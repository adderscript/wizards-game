import * as kaplay from 'kaplay';
import { makePlayer } from '../objects/player';
import { makeStaff } from '../objects/staff';
import { makeEnemyManager } from '../objects/enemyManager';

export function createGameScene(k: kaplay.KAPLAYCtx) {
    k.scene("game", () => {
        const player = makePlayer(k);
        const staff = makeStaff(k);
        const enemyManager = makeEnemyManager(k);

        k.add(player);
        player.add(staff);

        k.add(enemyManager);
    });
}
