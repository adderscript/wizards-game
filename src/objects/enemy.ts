import * as kaplay from 'kaplay';

export function makeEnemy(k: kaplay.KAPLAYCtx): kaplay.GameObj {
    const enemy = k.make([
        "enemy",
        k.sprite("enemy", { anim: "run", animSpeed: 0.8 }),
        k.pos(),
        k.scale(2.0),
        k.anchor("center"),
        k.area({ scale: 0.8 }),
        k.body({ drag: 12.5 }),
        k.health(2.0),

        {
            speed: 1000.0,

            damage: 1.0,
            knockbackForce: 1000.0,
        
            takeDamage(damage: number, knockbackForce: number, damagerPosition: kaplay.Vec2) {
                enemy.hurt(damage);

                // apply knockback
                const dir = enemy.pos.sub(damagerPosition).unit();
                enemy.applyImpulse(dir.scale(knockbackForce));
            },
        },
    ]);
 
    enemy.onUpdate(() => {
        const player = k.get("player")[0];
        if (player == null) {
            return;
        }

        // move towards player
        const dir = player.pos.sub(enemy.pos).unit();
        enemy.addForce(dir.scale(enemy.speed));
    });

    enemy.onCollide("player", (other: kaplay.GameObj) => {
        other.takeDamage(enemy.damage, enemy.knockbackForce, enemy.pos);
    });

    enemy.on("death", () => {
        enemy.destroy();
    });
    
    (enemy as any).takeDamage = (damage: number, knockbackForce: number, damagerPosition: kaplay.Vec2) => {
        enemy.hurt(damage);

        // apply knockback
        const dir = enemy.pos.sub(damagerPosition).unit();
        enemy.applyImpulse(dir.scale(knockbackForce));
    };

    return enemy;
}