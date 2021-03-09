
// The application will create a renderer using WebGL, if possible,
// with a fallback to a canvas render. It will also setup the ticker
// and the root stage PIXI.Container
let PixiApp = new PIXI.Application({ 
  width: 800,         // default: 800
  height: 600,        // default: 600
  antialias: true,    // default: false
  transparent: false, // default: false
  resolution: 1,       // default: 1
  backgroundColor: 0xd1e675
}
);

let PixiLoader = PIXI.Loader.shared;
let PixiResources = PixiLoader.resources;
let PixiSprite = PIXI.Sprite;


function pixiInit(){
// Insert the pixi cavas into the dom 
document.body.appendChild(PixiApp.view);

//make the pixi textures 
loadSprites();

//start the game
gameLoop();
}

let imageLocs = ["../images/sheep.png"];

function loadSprites(){
  //load an image and run the `setup` function when it's done
  PixiLoader
  .add(imageLocs)
  .load(addSprites);
}

function addSprites(){

  //Create the sheep sprite
  let sheep = new PixiSprite(PixiResources["../images/sheep.png"].texture);
  
  //Add the sheep to the stage
  PixiApp.stage.addChild(sheep);
}

function gameLoop(){
}
