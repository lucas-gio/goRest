use goRest;
db.getCollection("bicycles").deleteMany({});
db.getCollection("parts").deleteMany({});

db.getCollection("bicycles").insertMany([
    { code: "100", name:"Bicicleta Fiorenza 525 Cross", description:"BICICLETA RODADO 20 BICICLETA CROSS DE VARON NUEVO MODELO MARCA FIORENZA", origin: "ARG", unitPrice:"27199", image:"100" },
    { code: "101", name:"Bicicleta Kawasaki KHT 390", description:"Sku 513601 Bicicleta Kawasaki rodado 29 KHT 390, Cuadro de aluminio con horquilla de suspensión , Manubrio de aluminio en desnivel negro , manijas integradas ST-EF51 , Cambio 21 veloc. , Caja pedalera Cartridge , Frenos a disco , Llantas doble perfil de aluminio, origin Argentina , 6 meses de garantía., color negra con detalles verde y gris .", origin: "ARG", unitPrice:"87.299" , image:"101"},
    { code: "102", name:"Bicicleta Fiorenza 700 Beach", description:"Sku 503436 Bicicleta modelo 700 Beach Cruiser varón Rodado 26 , Llantas Aluminio, Rayos 2 mm, Cuadro acero Hi-Ten, Freno trasero contrapedal, asiento amplio, pie de apoyo, colores :naranja, azul ,amarillo ,negro,blanco,rojo", origin: "ARG", unitPrice:"28.199", image:"102" },
    { code: "103", name:"Bicicleta R 29 Daewoo Utah", description:"Sku 513997 Bicicleta R 29 marca Daewoo modelo Utah marco de acero , cambios 21 velocidades , manubrio y pedales aluminio, cadena YBN, frenos de disco, llantas aluminio doble pared, cubiertas nylon 2.1, piñon 7vel", origin: "ARG", unitPrice:"69.649" , image:"103"},
    { code: "104", name:"Bicicleta paseo Dama Fiorenza 464 Tempo Blanco y Verde", description:"BICICLETA DAMA RODADO 26 BICICLETA BICOLOR INCLUYE CANASTO CON GUARDABARRO CON PORTAPAQUETE ", origin: "ARG", unitPrice:"30.499" , image:"104"},
    { code: "105", name:"Bicicleta R 29 Euro bike Berlin", description:"bicicleta R 29 Erobike Acero, suspencion delantera 38MM aluminio ,manubrio y stem aluminio , cadena YBN , frenos y Shifter shimano EF500, frenos delantero y trasero disco , cubiertas nylon 2.1 llanatas aluminio , piñon 7 vel.", origin: "ARG", unitPrice:"70.299" , image:"105"},
    { code: "106", name:"Bicicleta Kawasaki modelo KST 210", description:"Bicleta R20 MTB Kawaasaki modelo KST 210 ,Cambios Shimano 7 Velocidades ,Frenos V-BRAKE de aluminio , horquilla aluminio,llantas de aluminio , cuadro de acero , ", origin: "ARG", unitPrice:"47.499" , image:"106"},
])

db.getCollection("parts").insertMany([
    { code: "100",
        name:"Wire Locks", description:"Wire Locks", origin: "ARG", unitPrice:"50000", image:"p1.jpg" },

    { code: "101", name:"Gloves",
        description:"Gloves", origin: "ARG", unitPrice:"40000", image:"p3.jpg" },

    { code: "102", name:"Cassete de velocidad",
        description:"Cassete de 6 velocidades nuevo", origin: "ARG", unitPrice:"20000", image:"p2.jpg" },

    { code: "103", name:"Pedales bicicleta de ruta",
        description:"Pedales de bicicleta de ruta, de acero inoxidable, rosca fina", origin: "ARG", unitPrice:"15000", image:"p4.jpg" },

    { code: "104", name:"Descarrilador trasero",
        description:"Descarrilador trasero, económico. ", origin: "ARG", unitPrice:"15000", image:"p5.JPG" },

    { code: "105", name:"Plato y palancas de carreras",
        description:"32 dientes, eje cuadrado", origin: "ARG", unitPrice:"7800", image:"p6.JPG" },

    { code: "106", name:"Manillar curvo",
        description:"Manillar cómodo, tamaño mediano", origin: "ARG", unitPrice:"6000", image:"p7.JPG" },

    { code: "107", name:"Cubiertas todo terreno",
        description:"Cubiertas negras, ideal para bicicletas de montaña", origin: "ARG", unitPrice:"9200", image:"p8.jpg" },

    { code: "108", name:"Conjunto de llantas",
        description:"Llantas trasera y delantera de aluminio", origin: "ARG", unitPrice:"50682", image:"p9.jpg" },

    { code: "109", name:"Plato de acero inoxidable",
        description:"38 dientes", origin: "ARG", unitPrice:"5800", image:"p10.jpg" },

    { code: "110", name:"Suspención delantera",
        description:"Regulable. Colores diponibles: cromado, negro mate, negro brillante.",
        origin: "ARG", unitPrice:"6800", image:"p11.jpg" },

    { code: "111", name:"Casco plástico reforzado",
        description:"Casco reforzado ajustable. Colores negro, rojo, blanco, amarillo.",
        origin: "ARG", unitPrice:"1500", image:"p12.jpg" },

    { code: "112", name:"Marco medio. Ideal para crear tu bicicleta.",
        description:"Marco de acero inoxidable. Reforzado. Modelos MTB y Playera.",
        origin: "ARG", unitPrice:"1500", image:"p12.jpg" },

    { code: "113", name:"Asiento confort",
        description:"Asiento cómodo para playera. Colores negro y gris disponibles.",
        origin: "ARG", unitPrice:"3500", image:"p14.jpg" },

    { code: "114", name:"Rodillera motocross",
        description:"Rodillera plástica ajustable.",
        origin: "ARG", unitPrice:"1500", image:"p15.jpg" },

    { code: "115", name:"Inflador de pie",
        description:"Tamaño medio",
        origin: "ARG", unitPrice:"8000", image:"p16.JPG" },
])