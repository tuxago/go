package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/window"
)

var title *gui.Label

// Deserialize JPlayers
type JPlayers struct {
	Players []JPlayer
}

type JPlayer struct {
	Name string
	Wins int
}

var numbWins = 0

//go:embed earth.vert
var shaderEarthVertex string

//go:embed earth.frag
var shaderEarthFrag string

func main() {
	a := app.App()
	scene := core.NewNode()
	e := Earth{
		app:   a,
		scene: scene,
	}
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	e.setupGUI()

	e.start()
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
		e.mainPanel.SetSize(float32(width), float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)
	go func() {
		//update player wins from server http://localhost:8080/wins
		for range time.Tick(1 * time.Second) {
			wins, err := http.Get("http://localhost:8080/players")
			if err != nil {
				log.Println(err)
				//avoid crash
				continue
			}
			defer wins.Body.Close()
			var players JPlayers
			err = json.NewDecoder(wins.Body).Decode(&players)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, player := range players.Players {
				if player.Name == "John" {
					numbWins = player.Wins
				}
			}
			title.SetText(fmt.Sprintf("WINS: %d", numbWins))
			// wins.Body.Close()
		}
	}()

	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
		e.Update(deltaTime)
	})
}

func (e *Earth) setupGUI() {
	dl := gui.NewDockLayout()
	width, height := e.app.GetSize()
	e.mainPanel = gui.NewPanel(float32(width), float32(height))
	e.mainPanel.SetRenderable(false)
	e.mainPanel.SetEnabled(false)
	e.mainPanel.SetLayout(dl)
	e.scene.Add(e.mainPanel)
	gui.Manager().Set(e.mainPanel)

	headerColor := math32.Color4{R: 13.0 / 256.0, G: 41.0 / 256.0, B: 62.0 / 256.0, A: 1}
	lightTextColor := math32.Color4{R: 0.8, G: 0.8, B: 0.8, A: 1}
	header := gui.NewPanel(600, 40)
	header.SetBorders(0, 0, 1, 0)
	header.SetPaddings(4, 4, 4, 4)
	header.SetColor4(&headerColor)
	header.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})

	// Horizontal box layout for the header
	hbox := gui.NewHBoxLayout()
	header.SetLayout(hbox)
	e.mainPanel.Add(header)

	// Header title
	const fontSize = 50
	title = gui.NewLabel(" ")
	title.SetFontSize(fontSize)
	title.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	title.SetText(fmt.Sprintf("WINS: %d", numbWins))
	title.SetColor4(&lightTextColor)
	header.Add(title)
}

type Earth struct {
	app *app.Application

	mainPanel *gui.Panel

	sphere *graphic.Mesh
	scene  *core.Node
}

func (e *Earth) start() {
	// Create Skybox
	skybox, err := graphic.NewSkybox(graphic.SkyboxData{
		DirAndPrefix: "./images/space/dark-s_", Extension: "jpg",
		Suffixes: [6]string{"px", "nx", "py", "ny", "pz", "nz"}})
	if err != nil {
		panic(err)
	}
	e.scene.Add(skybox)

	// Adds directional front light
	dir1 := light.NewDirectional(&math32.Color{R: 1, G: 1, B: 1}, 0.9)
	dir1.SetPosition(0, 0, 100)
	e.scene.Add(dir1)

	// Create custom shader
	e.app.Renderer().AddShader("shaderEarthVertex", shaderEarthVertex)
	e.app.Renderer().AddShader("shaderEarthFrag", shaderEarthFrag)
	e.app.Renderer().AddProgram("shaderEarth", "shaderEarthVertex", "shaderEarthFrag")

	// Helper function to load a texture and handle errors
	newTexture := func(path string) *texture.Texture2D {
		tex, err := texture.NewTexture2DFromImage(path)
		if err != nil {
			log.Fatalf("Error loading texture: %s", err)
		}
		tex.SetFlipY(false)
		return tex
	}

	// Create earth textures
	texDay := newTexture("./images/earth_clouds_big.jpg")
	texSpecular := newTexture("./images/earth_spec_big.jpg")
	texNight := newTexture("./images/earth_night_big.jpg")
	//texBump, err := newTexture("./images/earth_bump_big.jpg")

	// Create custom material using the custom shader
	matEarth := NewEarthMaterial(&math32.Color{R: 1, G: 1, B: 1})
	matEarth.SetShininess(20)
	//matEarth.SetSpecularColor(&math32.Color{0., 1, 1})
	//matEarth.SetColor(&math32.Color{0.8, 0.8, 0.8})
	matEarth.AddTexture(texDay)
	matEarth.AddTexture(texSpecular)
	matEarth.AddTexture(texNight)

	// Create sphere
	geom := geometry.NewSphere(1, 32, 16)
	e.sphere = graphic.NewMesh(geom, matEarth)
	e.scene.Add(e.sphere)

	// Create sun sprite
	texSun, err := texture.NewTexture2DFromImage("./images/lensflare0_alpha.png")
	if err != nil {
		log.Fatalf("Error loading texture: %s", err)
	}
	sunMat := material.NewStandard(&math32.Color{R: 1, G: 1, B: 1})
	sunMat.AddTexture(texSun)
	sunMat.SetTransparent(true)
	sun := graphic.NewSprite(10, 10, sunMat)
	sun.SetPositionZ(20)
	e.scene.Add(sun)

	// Add axes helper
	//axes := helper.NewAxes(5)
	//e.scene.Add(axes)
}

// Update is called every frame.
func (t *Earth) Update(deltaTime time.Duration) {
	t.sphere.RotateY(0.1 * float32(deltaTime.Seconds()))
}

type EarthMaterial struct {
	material.Standard // Embedded standard material
}

// NewEarthMaterial creates and returns a pointer to a new earth material
func NewEarthMaterial(color *math32.Color) *EarthMaterial {

	pm := new(EarthMaterial)
	pm.Standard.Init("shaderEarth", color)
	return pm
}
