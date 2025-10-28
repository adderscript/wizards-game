import * as kaplay from 'kaplay';

export function makePlayer(k: kaplay.KAPLAYCtx): kaplay.GameObj {
    const player = k.make([
        "player",
        k.sprite("player", { anim: "idle", animSpeed: 0.6, }),
        k.pos(k.center()),
        k.scale(2.0),
        k.anchor("center"),
        k.area({ scale: 0.4 }),
        k.body({ drag: 12.5 }),
        k.health(3.0),
        
        {
            speed: 100000.0,
        },
    ]);

    player.onUpdate(() => {
        const isMoving = player.vel.len() > 50;
        const currentAnim = player.getCurAnim()?.name; 

        // play animations
        if (isMoving && currentAnim !== "run") {
            player.play("run");
        } else if (!isMoving && currentAnim !== "idle") {
            player.play("idle");
        }
    });

    player.onKeyDown("d", () => {
        player.addForce(k.vec2(1.0, 0.0).scale(player.speed * k.dt()));
    });
    player.onKeyDown("a", () => {
        player.addForce(k.vec2(-1.0, 0.0).scale(player.speed * k.dt()));
    });
    player.onKeyDown("s", () => {
        player.addForce(k.vec2(0.0, 1.0).scale(player.speed * k.dt()));
    });
    player.onKeyDown("w", () => {
        player.addForce(k.vec2(0.0, -1.0).scale(player.speed * k.dt()));
    });

    player.on("death", () => {
        player.destroy();
    });

    (player as any).takeDamage = (damage: number, knockbackForce: number, damagerPosition: kaplay.Vec2) => {
        player.hurt(damage);

        // apply knockback
        const dir = player.pos.sub(damagerPosition).unit();
        player.applyImpulse(dir.scale(knockbackForce));
        k.shake(2.0);
    };

    return player;
}