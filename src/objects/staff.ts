import * as kaplay from 'kaplay';
import { makeBullet } from './bullet';

export function makeStaff(k: kaplay.KAPLAYCtx): kaplay.GameObj {
    const staff = k.make([
        "staff",
        k.sprite("staff"),
        k.pos(),
        k.rotate(),
        k.anchor("center"),

        {
            distance: 10.0,
            shootForce: 250.0,
        },
    ]);

    const firepoint = staff.add([
        k.pos(1.0, -1.0),
    ])

    staff.onUpdate(() => {
        const player = k.get("player")[0];

        // point towards mouse
        const diff = k.mousePos().sub(player.pos);
        staff.rotateTo(diff.angle());

        // move away from player
        const dir = k.vec2(Math.cos(k.deg2rad(staff.angle)), Math.sin(k.deg2rad(staff.angle)));
        staff.moveTo(dir.scale(staff.distance));
    });

    staff.onMousePress("left", () => {
        const bullet = makeBullet(k);

        // shoot in facing direction
        const dir = k.vec2(Math.cos(k.deg2rad(staff.angle)), Math.sin(k.deg2rad(staff.angle)));
        bullet.applyImpulse(dir.scale(staff.shootForce));
        bullet.moveTo(firepoint.worldPos());
        bullet.rotateTo(staff.angle);

        k.add(bullet);
    })

    return staff;
}