// Colección Héroes

db.heroes.insertMany([
    {
        urlImagen: '',
        clase: 'Guerrero',
        tipo: 'Tanque',
        poder: 1,
        vida: 24,
        defensa: 11,
        ataqueBase: 10,
        ataqueDado: 6,
        danoMax: 4,
        activo: true
    },
    {
        urlImagen: '',
        clase: 'Guerrero',
        tipo: 'Armas',
        poder: 1,
        vida: 24,
        defensa: 11,
        ataqueBase: 10,
        ataqueDado: 6,
        danoMax: 4,
        activo: true
    },
    {
        urlImagen: '',
        clase: 'Mago',
        tipo: 'Fuego',
        poder: 1,
        vida: 20,
        defensa: 10,
        ataqueBase: 10,
        ataqueDado: 8,
        danoMax: 6,
        activo: true
    },
    {
        urlImagen: '',
        clase: 'Mago',
        tipo: 'Hielo',
        poder: 1,
        vida: 20,
        defensa: 10,
        ataqueBase: 10,
        ataqueDado: 8,
        danoMax: 6,
        activo: true
    },
    {
        urlImagen: '',
        clase: 'Pícaro',
        tipo: 'Veneno',
        poder: 1,
        vida: 16,
        defensa: 8,
        ataqueBase: 10,
        ataqueDado: 10,
        danoMax: 8,
        activo: true
    },
    {
        urlImagen: '',
        clase: 'Pícaro',
        tipo: 'Machete',
        poder: 1,
        vida: 16,
        defensa: 8,
        ataqueBase: 10,
        ataqueDado: 10,
        danoMax: 8,
        activo: true
    },
]);

/* Colección Armas

statEffect: Buff o debuff que recibirá la stat
stat: Estadística que afecta
target: A quién apunta el tipo de efecto (Player o Enemy, jugador o enemigo)
turnCount: Por cuántos turnos se juega la carta (los efectos son permanentes)

case 0: Suma o resta una estadística única por turnos determinados, a un objetivo.
case 1: Compara dos estadísticas, y está a favor de una para sumar o restar una estadística única por turnos determinados, a un objetivo [Stat a favor, stat 0, stat 1, stat afectada].
case 2: Afecta a dos estadísticas, sumando o restando por turnos determinados, a un objetivo.
case 3: Solo toma efecto cuando se ha utilizado otra carta, modifica el efecto de la carta anterior.
case 4: Una estadística se multiplica por un valor específico durante un turno.
*/

db.armas.insertMany([
    {
        urlImagen: '',
        nombre: 'Espada de una mano',
        tipoHeroe: 'Guerrero Tanque',
        efecto: { case: 0, statEffect: 1, stat: 'Ataque', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Escudo de dragón',
        tipoHeroe: 'Guerrero Tanque',
        efecto: { case: 0, statEffect: 1, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Espada de dos manos',
        tipoHeroe: 'Guerrero Armas',
        efecto: { case: 0, statEffect: 1, stat: 'Ataque', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Piedra de afilar',
        tipoHeroe: 'Guerrero Armas',
        efecto: { case: 0, statEffect: 2, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Orbe de manos ardientes',
        tipoHeroe: 'Mago Fuego',
        efecto: { case: 0, statEffect: 1, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Fuego fatuo',
        tipoHeroe: 'Mago Fuego',
        efecto: { case: 0, statEffect: 1, stat: 'Ataque', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Báculo de Permafrost',
        tipoHeroe: 'Mago Hielo',
        efecto: { case: 0, statEffect: -1, stat: 'Daño', target: 'Enemy', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Venas heladas',
        tipoHeroe: 'Mago Hielo',
        efecto: { case: 0, statEffect: 1, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Daga purulenta',
        tipoHeroe: 'Pícaro Veneno',
        efecto: { case: 0, statEffect: 1, stat: 'Daño', target: 'Player', turnCount: 2 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Visión borrosa',
        tipoHeroe: 'Pícaro Veneno',
        efecto: { case: 0, statEffect: -1, stat: 'Ataque', target: 'Enemy', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Machete bendito',
        tipoHeroe: 'Pícaro Machete',
        efecto: { case: 0, statEffect: 2, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        nombre: 'Sierra sangrienta',
        tipoHeroe: 'Pícaro Machete',
        efecto: { case: 0, statEffect: 2, stat: 'Daño', target: 'Player', turnCount: 2 },
        activo: true
    }
]);

// Colección Armaduras

db.armaduras.insertMany([
    {
        urlImagen: '',
        heroe: 'Guerrero Tanque',
        tipo: 'Pesada',
        efecto: { case: 0, statEffect: 4, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Guerrero Armas',
        tipo: 'Media',
        efecto: { case: 0, statEffect: 3, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Fuego',
        tipo: 'Tela',
        efecto: { case: 0, statEffect: 1, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Hielo',
        tipo: 'Tela',
        efecto: { case: 0, statEffect: 1, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Veneno',
        tipo: 'Media liviana',
        efecto: { case: 0, statEffect: 2, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Machete',
        tipo: 'Media liviana',
        efecto: { case: 0, statEffect: 2, stat: 'Defensa', target: 'Player', turnCount: 1 },
        activo: true
    }
]);

// Colección Ítems
// ARREGLAR (qué estadística es afectada?) Pinchos de escudo
// PENDIENTE Veneno lacerante, Mancuerna yugular

db.items.insertMany([
    {
        urlImagen: '',
        heroe: 'Guerrero Tanque',
        nombre: 'Pinchos de escudo',
        efecto: { case: 1, statEffect: -1, stat: ['1', 'Ataque', 'Defensa', 'Vida'], target: 'Enemy', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Guerrero Armas',
        nombre: 'Empuñadura de Furia',
        efecto: { case: 2, statEffect: [1, -1], stat: ['Daño', 'Vida'], target: 'Player', turnCount: 2 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Fuego',
        nombre: 'Anillo para Piro-Explosión',
        efecto: { case: 0, statEffect: 3, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Hielo',
        nombre: 'Libro de la ventisca helada',
        efecto: { case: 0, statEffect: 2, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Veneno',
        nombre: 'Veneno lacerante',
        efecto: '',
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Machete',
        nombre: 'Mancuerna yugular',
        efecto: { case: 3 },
        activo: true
    }
]);

// Todas las cartas son exclusivas del héroe?

db.habilidades.insertMany([
    {
        urlImagen: '',
        heroe: 'Guerrero Tanque',
        nombre: 'Golpe de defensa',
        efectoGlobal: { case: 0, statEffect: 1, stat: 'Ataque', target: 'Player', turnCount: 1 },
        efectoHeroe: { case: 0, statEffect: 4, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Guerrero Armas',
        nombre: 'Segundo impulso',
        efectoGlobal: { case: 0, statEffect: [1, 2, 3, 4], stat: 'Vida', target: 'Player', turnCount: 1 },
        efectoHeroe: { case: 0, statEffect: 3, stat: 'Vida', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Fuego',
        nombre: 'Luz cegadora',
        efectoGlobal: { case: 0, statEffect: 1, stat: 'Vida', target: 'Player', turnCount: 1 },
        efectoHeroe: { case: 4, statEffect: 0, stat: 'Daño', target: 'Enemy', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Mago Hielo',
        nombre: 'Frio concentrado',
        efectoGlobal: { case: 0, statEffect: -1, stat: 'Poder', target: 'Enemy', turnCount: 1 },
        efectoHeroe: { case: 0, statEffect: 2, stat: 'Daño', target: 'Player', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Veneno',
        nombre: 'Toma y lleva',
        efectoGlobal: { case: 0, statEffect: 1, stat: 'Ataque', target: 'Player', turnCount: 1 },
        efectoHeroe: { case: 4, statEffect: 0.5, stat: 'Daño', target: 'Enemy', turnCount: 1 },
        activo: true
    },
    {
        urlImagen: '',
        heroe: 'Pícaro Machete',
        nombre: 'Intimidación sangrienta',
        efectoGlobal: { case: 0, statEffect: 1, stat: 'Daño', target: 'Player', turnCount: 1 },
        efectoHeroe: { case: 0, statEffect: 2, stat: 'Vida', target: 'Player', turnCount: 1 },
        activo: true
    }
]);