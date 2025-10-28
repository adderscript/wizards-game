import * as kaplay from 'kaplay';

export function makeBullet(k: kaplay.KAPLAYCtx): kaplay.GameObj {
    const bullet = k.make([
        "bullet",
        k.sprite("bullet", { anim: "idle" }),
        k.pos(),
        k.scale(2.0),
        k.rotate(),
        k.anchor("center"),
        k.area({ scale: 0.6 }),
        k.body(),
        k.offscreen({ destroy: true }),

        {
            damage: 1.0,
            knockbackForce: 500.0,
        },
    ]);

    bullet.onCollide("enemy", (other: kaplay.GameObj) => {
        other.takeDamage(bullet.damage, bullet.knockbackForce, bullet.pos);
        bullet.destroy();
    });

    return bullet;
}