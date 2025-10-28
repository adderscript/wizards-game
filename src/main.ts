import kaplay from 'kaplay';
import { createGameScene } from './scenes/game';

const k = kaplay({
    background: [0.0, 0.0, 0.0],
    width: 640,
    height: 360,
    stretch: false,
    letterbox: true,
    crisp: true,
    debug: true,
    global: false,
});

k.loadAseprite("player", "sprites/player/player.png", "sprites/player/player.json");
k.loadSprite("staff", "sprites/staff.png");
k.loadAseprite("bullet", "sprites/bullet/bullet.png", "sprites/bullet/bullet.json");
k.loadAseprite("enemy", "sprites/enemy/enemy.png", "sprites/enemy/enemy.json");

createGameScene(k);
k.go("game");