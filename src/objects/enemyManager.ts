import * as kaplay from 'kaplay';
import { makeEnemy } from './enemy';

export function makeEnemyManager(k: kaplay.KAPLAYCtx): kaplay.GameObj {
    const enemyManager = k.make([
        "enemyManager",
        k.timer(),
        
        {
            spawnDelay: 2.0,
        },
    ]);

    enemyManager.loop(enemyManager.spawnDelay, () => {
        // spawn enemy
        const enemy = makeEnemy(k);
        k.add(enemy);
    });

    return enemyManager;
}