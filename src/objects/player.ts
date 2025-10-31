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

            invincibilityDelay: 0.5,
            invincibilityTimer: 0.0,

            takeDamage(damage: number, knockbackForce: number, damagerPosition: kaplay.Vec2) {
                if (player.invincibilityTimer < player.invincibilityDelay) {
                    return;
                }

                player.hurt(damage);

                // apply knockback
                const dir = player.pos.sub(damagerPosition).unit();
                player.applyImpulse(dir.scale(knockbackForce));
                k.shake(2.0);

                // start invincibility
                player.invincibilityTimer = 0.0;
            },
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

        // update invincibility timer
        if (player.invincibilityTimer < player.invincibilityDelay) {
            player.invincibilityTimer += k.dt();
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

    return player;
}