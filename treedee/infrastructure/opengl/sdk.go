package opengl

import (
	"time"

	"github.com/deepvalue-network/software/treedee/application/windows"
	domain_window "github.com/deepvalue-network/software/treedee/domain/windows"
	"github.com/deepvalue-network/software/treedee/domain/worlds"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/displays"
	hud_nodes "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/nodes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries"
	vertex_shaders "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures"
	texture_shaders "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/shaders"
	fragment_shader "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
	"github.com/go-gl/mathgl/mgl32"
	uuid "github.com/satori/go.uuid"
)

const maxAlpha = 0xff

const float32SizeInBytes = 32 / 8

const glStrPattern = "%s\x00"

// WorldLogicFn represents the world logic func
type WorldLogicFn func(world World) error

// WorldBuilder represents a world builder
type WorldBuilder interface {
	Create() WorldBuilder
	WithWindow(win domain_window.Window) WorldBuilder
	WithWorld(world worlds.World) WorldBuilder
	WithLogic(fn WorldLogicFn) WorldBuilder
	Now() (World, error)
}

// World represents a world index
type World interface {
	ID() *uuid.UUID
	WindowApplication() windows.Application
	Logic() WorldLogicFn
	CurrentSceneIndex() uint
	Scenes() []Scene
	Execute() error
}

// SceneBuilder represents a scene builder
type SceneBuilder interface {
	Create() SceneBuilder
	WithScene(scene scenes.Scene) SceneBuilder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	ID() *uuid.UUID
	Index() uint
	Hud() Hud
	Nodes() []Node
	Render(
		delta time.Duration,
		activeCameraID *uuid.UUID,
	) error
}

// HudBuilder represents a hud builder
type HudBuilder interface {
	Create() HudBuilder
	WithHud(hud huds.Hud) HudBuilder
	Now() (Hud, error)
}

// Hud represents a head-up display
type Hud interface {
	ID() *uuid.UUID
	Program() uint32
	HasNodes() bool
	Nodes() []HudNode
	HasMaterial() bool
	Material() Material
	Render(
		delta time.Duration,
		activeScene Scene,
	) error
}

// HudNodeBuilder represents a hud node builder
type HudNodeBuilder interface {
	Create() HudNodeBuilder
	WithNode(node hud_nodes.Node) HudNodeBuilder
	Now() (HudNode, error)
}

// HudNode represents a hud node
type HudNode interface {
	ID() *uuid.UUID
	Position() HudPosition
	Orientation() Orientation
	HasContent() bool
	Content() HudNodeContent
	HasChildren() bool
	Children() []HudNode
	Render(
		delta time.Duration,
		activeScene Scene,
	) error
}

// HudNodeContent represents the hud node content
type HudNodeContent interface {
	IsModel() bool
	Model() Model
	IsDisplay() bool
	Display() HudDisplay
}

// HudPosition represents a hud position
type HudPosition interface {
	Vector() mgl32.Vec2
	Variable() string
	Add(pos HudPosition) HudPosition
}

// HudDisplayBuilder represents a hud display builder
type HudDisplayBuilder interface {
	Create() HudDisplayBuilder
	WithDisplay(display displays.Display) HudDisplayBuilder
	Now() (HudDisplay, error)
}

// HudDisplay represents a hud display
type HudDisplay interface {
	ID() *uuid.UUID
	Index() uint
	Viewport() viewports.Viewport
	Camera() Camera
	HasMaterial() bool
	Material() Material
}

// NodeBuilder represents a node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithNode(node nodes.Node) NodeBuilder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	ID() *uuid.UUID
	Position() Position
	Orientation() Orientation
	HasContent() bool
	Content() NodeContent
	HasChildren() bool
	Children() []Node
	Render(
		delta time.Duration,
		activeCamera WorldCamera,
		activeScene Scene,
	) error
}

// NodeContent represents the node content
type NodeContent interface {
	IsModel() bool
	Model() Model
	IsCamera() bool
	Camera() Camera
}

// CameraBuilder represents the camera builder
type CameraBuilder interface {
	Create() CameraBuilder
	WithCamera(cam cameras.Camera) CameraBuilder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	ID() *uuid.UUID
	Index() uint
	Projection() CameraProjection
	LookAt() CameraLookAt
	Render(
		delta time.Duration,
		pos Position,
		orientation Orientation,
		activeScene Scene,
		program uint32,
	) error
}

// CameraLookAt represents the direction where the camera looks at
type CameraLookAt interface {
	Variable() string
	Eye() mgl32.Vec3
	Center() mgl32.Vec3
	Up() mgl32.Vec3
}

// CameraProjection represents the camera projection
type CameraProjection interface {
	Variable() string
	FieldOfView() float32
	AspectRation() float32
	Near() float32
	Far() float32
}

// ModelBuilder represents a model builder
type ModelBuilder interface {
	Create() ModelBuilder
	WithModel(model models.Model) ModelBuilder
	Now() (Model, error)
}

// Model represents a model
type Model interface {
	ID() *uuid.UUID
	Program() uint32
	Geometry() Geometry
	Material() Material
	Render(
		delta time.Duration,
		activeCamera WorldCamera,
		activeScene Scene,
	) error
}

// GeometryBuilder represents a geometry builder
type GeometryBuilder interface {
	Create() GeometryBuilder
	WithProgram(prog uint32) GeometryBuilder
	WithGeometry(geo geometries.Geometry) GeometryBuilder
	Now() (Geometry, error)
}

// Geometry represents a geometry
type Geometry interface {
	ID() *uuid.UUID
	Program() uint32
	VAO() uint32
	VertexAmount() int32
	VertexType() VertexType
	Shader() VertexShader
	Prepare() error
	Render() error
}

// VertexType represents vertex type
type VertexType interface {
	IsTriangle() bool
}

// VertexShader represents a vertex shader
type VertexShader interface {
	ID() *uuid.UUID
	Variables() VertexShaderVariables
}

// VertexShaderVariables represents vertex shader variables
type VertexShaderVariables interface {
	TextureCoordinates() string
	VertexCoordinates() string
}

// MaterialBuilder represents a material builder
type MaterialBuilder interface {
	Create() MaterialBuilder
	WithMaterial(mat materials.Material) MaterialBuilder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	ID() *uuid.UUID
	Alpha() Alpha
	Viewport() viewports.Viewport
	Layers() []Layer
	Render(
		delta time.Duration,
		pos Position,
		orientation Orientation,
		activeScene Scene,
		program uint32,
	) error
}

// LayerBuilder represents a layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithLayer(layer layers.Layer) LayerBuilder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	ID() *uuid.UUID
	Index() uint
	Alpha() Alpha
	Texture() Texture
	Render(
		delta time.Duration,
		pos Position,
		orientation Orientation,
		activeScene Scene,
		program uint32,
	) error
}

// TextureBuilder represents a texture builder
type TextureBuilder interface {
	Create() TextureBuilder
	WithTexture(tex textures.Texture) TextureBuilder
	Now() (Texture, error)
}

// Texture represents a texture
type Texture interface {
	ID() *uuid.UUID
	Dimension() ints.Vec2
	Variable() string
	IsResource() bool
	Resource() uint32
	IsCamera() bool
	Camera() Camera
	IsShader() bool
	Shader() TextureShader
	Render(
		delta time.Duration,
		pos Position,
		orientation Orientation,
		activeScene Scene,
		program uint32,
	) error
}

// TextureShader represents a texture shader
type TextureShader interface {
	ID() *uuid.UUID
	Program() uint32
	IsDynamic() bool
	Render(
		delta time.Duration,
		pos Position,
		orientation Orientation,
		program uint32,
	) error
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithTextureShader(texShader texture_shaders.Shader) ProgramBuilder
	WithVertexShader(verShader vertex_shaders.Shader) ProgramBuilder
	WithFragmentShader(fragShader fragment_shader.Shader) ProgramBuilder
	Now() (uint32, error)
}

// Alpha returns the alpha
type Alpha interface {
	Value() float32
	Variable() string
}

// Position represents a 3d position
type Position interface {
	Vector() mgl32.Vec3
	Variable() string
	Add(pos Position) Position
}

// Orientation represents an orientation
type Orientation interface {
	Angle() float32
	Direction() mgl32.Vec3
	Variable() string
	Add(orientation Orientation) Orientation
}

// WorldCameraBuilder represents a world camera builder
type WorldCameraBuilder interface {
	Create() WorldCameraBuilder
	WithNode(node Node) WorldCameraBuilder
	Now() (WorldCamera, error)
}

// WorldCamera represents a camera in a 3d world
type WorldCamera interface {
	Camera() Camera
	Position() Position
	Orientation() Orientation
	Update(pos Position, orientation Orientation) WorldCamera
	Slide(pos HudPosition, orientation Orientation) WorldCamera
}
