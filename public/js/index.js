
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

let FPS = {
  CheckPointTimeStamp : Date.now(),
  Count : 0,
  Log : [],
  CurrentSmoothed : 60,
  Show: true,
  Text:{}, // <- will be set in Init
  Init: function(app, shouldShow){
    this.Show = shouldShow;
    let style = new PIXI.TextStyle({
      fill: "#333333",
      fontSize: 40,
      fontWeight: 'bold',
    });

    this.Text = new PIXI.Text();
    this.Text.anchor.set(.5); //Center text to anchor
    this.Text.x = 775;
    this.Text.y = 575;

    app.stage.addChild(this.Text);
  },
  Tick: function(){
    let now = Date.now();
    this.Count++;
    
    if (now - this.CheckPointTimeStamp > 1000){
      this.CheckPointTimeStamp = now;
      this.Log.push(this.Count);
      this.Count = 0;
      this.CurrentSmoothed = this.Log.reduce((a, b) => a + b, 0) / this.Log.length;
      if (this.Log.length > 10){
        this.Log.shift();
      }
    }
    if(this.Show){
      let currentFrames = Math.floor(this.CurrentSmoothed);
      //console.log(currentFrames);
      this.Text.text = currentFrames;
    }
  }
}


function pixiInit(){
// Insert the pixi cavas into the dom 
document.body.appendChild(PixiApp.view);
FPS.Init(PixiApp, true);


//make the pixi textures 
loadSpriteTextures();

//start the game
requestAnimationFrame(gameLoop);
}

let imageLocs = ["../images/sheep.png"];

function loadSpriteTextures(){
  //load an image and run the `setup` function when it's done
  PixiLoader
  .add(imageLocs)
  .load(spritesLoaded);
}

function spritesLoaded(){

  //Create the sheep sprite
  let sheep = new PixiSprite(PixiResources["../images/sheep.png"].texture);
  
  //Add the sheep to the stage
  PixiApp.stage.addChild(sheep);
}

function gameLoop(){
  FPS.Tick();
  requestAnimationFrame(gameLoop);
}


function handleMessage(msg){

}

