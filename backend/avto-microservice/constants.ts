const auto1 = '/uploads/cars/17.png'
const auto2 = '/uploads/cars/20.png'
const auto3 = '/uploads/cars/22.png'
const auto4 = '/uploads/cars/26.png'
const auto5 = '/uploads/cars/30.png'
const auto6 = '/uploads/cars/34.png'
const auto7 = '/uploads/cars/37.png'
const auto8 = '/uploads/cars/45.png'

export type TCar = {
  img: string //Изображение автовышки
  imgArray: string[] //Изображения автовышки для галереи
  height: number //Высота автовышки
  type: string //Тип автовышки
  power: number //Грузоподъемность автовышки
  price5: number //Цена 5 НДС
  price22: number //Цена 22 НДС
  description: string //Описание автовышки
  brand: string //Марка автовышки
  machine: string //Установка автовышки
  sizeTs: string[] //Габариты в трансполртном положении (длина, ширина, высота)
  width: string //Ширина с опорами
  mass: string //Масса, тонн
  sizeCradle: string[] | string[][] //Габариты люльки, ширина (в сложенном и разложенном состоянии) и длина (в сложенном и разложенном состоянии)
  special: string //Особенности автовышки
  rostechReg: boolean //Регистрации в ростехнадзоре
}

export const cars: TCar[] = [
  {
    img: auto1,
    imgArray: [],
    height: 17,
    type: 'Телескопическая',
    power: 800,
    price5: 2600,
    price22: 3050,
    description: 'Платформа 2х4, съемные борта, грузоподъемность до 800 кг',
    brand: 'MITSUBISHI CANTER, ISUZU ELF',
    machine: 'AICHI, TADANO',
    sizeTs: ['6.40', '2.00', '3.60'],
    width: '4.30',
    mass: '7.50',
    sizeCradle: ['4.00', '1.90'],
    special: 'Люлька «балкон», съемные борта',
    rostechReg: true,
  },
  {
    img: auto2,
    imgArray: [],
    height: 20,
    type: 'Телескопическая',
    power: 800,
    price5: 3000,
    price22: 3500,
    description: 'Платформа 2х4, съемные борта, грузоподъемность до 800 кг',
    brand: 'MITSUBISHI FUSO FIGHTER, ISUZU FORWARD',
    machine: 'TADANO',
    sizeTs: ['6.90', '2.10', '3.40'],
    width: '5.00',
    mass: '8.00',
    sizeCradle: [
      ['3.00', '2.00'],
      ['4.00', '2.00'],
    ],
    special: 'Люлька «балкон», съемные борта',
    rostechReg: true,
  },
  {
    img: auto3,
    imgArray: [],
    height: 22,
    type: 'Телескопическая',
    power: 800,
    price5: 3300,
    price22: 3850,
    description: 'Платформа 2х4, съемные борта, грузоподъемность до 800 кг',
    brand: 'ISUZU ELF',
    machine: 'TADANO',
    sizeTs: ['7.20', '2.20', '3.60'],
    width: '4.60',
    mass: '8.50',
    sizeCradle: ['4.00', '1.90'],
    special: 'Люлька «балкон», съемные борта',
    rostechReg: true,
  },
  {
    img: auto4,
    imgArray: [],
    height: 26,
    type: 'Телескопическая',
    power: 300,
    price5: 3400,
    price22: 3950,
    description:
      'Люлька 1х1,5, раскладывается до размеров 1х3 возможно управление сверху (из люльки), грузоподъемность до 200 кг',
    brand: 'ISUZU FORWARD',
    machine: 'AICHI',
    sizeTs: ['7.00', '2.00', '3.50'],
    width: '4.40',
    mass: '8.50',
    sizeCradle: [
      ['0.80', '1.50'],
      ['1.00', '3.00'],
    ],
    special: 'Люлька поворотная, раскладывается',
    rostechReg: true,
  },
  {
    img: auto5,
    imgArray: [],
    height: 30,
    type: 'Телескопическая',
    power: 300,
    price5: 4200,
    price22: 4900,
    description:
      'Есть лебедка грузоподъемностью 1тн, люлька раскладывается до размеров 1х3, грузоподъемность до 300 кг',
    brand: 'MITSUBISHI FUSO CANTER',
    machine: 'TADANO',
    sizeTs: ['9.70', '2.20', '3.60'],
    width: '5.50',
    mass: '9.00',
    sizeCradle: ['1.00', '0.7'],
    special: 'Есть управление в люльке',
    rostechReg: true,
  },
  {
    img: auto6,
    imgArray: [],
    height: 34,
    type: 'Телескоп + колено',
    power: 300,
    price5: 5300,
    price22: 6150,
    description:
      'Телескоп с коленом, может подавать на высоту: -5 м, поворотная люлька 1х1,5, возможно управление сверху, грузоподъемность до 300 кг',
    brand: 'КАМАЗ-53215',
    machine: 'ПСС-141.32',
    sizeTs: ['9.00', '2.50', '3.40'],
    width: '5.70',
    mass: '19.00',
    sizeCradle: ['1.60', '1.00'],
    special: 'Телескоп с коленом, может подавать на высоту: -5 м, возможно управление сверху (из люльки)',
    rostechReg: true,
  },
  {
    img: auto7,
    imgArray: [],
    height: 37,
    type: 'Телескоп + колено',
    power: 300,
    price5: 5800,
    price22: 6750,
    description:
      'Характеристики: Вездеход, телескоп с коленом, люлька 1х1,5, возможно управление сверху, грузоподъемность до 300 кг',
    brand: 'КАМАЗ-43118',
    machine: 'ПСС-141.32',
    sizeTs: ['9.00', '2.50', '3.70'],
    width: '5.70',
    mass: '19.50',
    sizeCradle: ['1.60', '1.00'],
    special: 'Вездеход, телескоп с коленом, может подавать на высоту: -5 м, возможно управление сверху (из люльки)',
    rostechReg: true,
  },
  {
    img: auto8,
    imgArray: [],
    height: 45,
    type: 'Телескоп + стрела и рукоять',
    power: 450,
    price5: 7500,
    price22: 6750,
    description:
      ' Двухсекционная рукоять стрелы, с возможностью выдвижения на 3,5 метра, дает возможность работать ниже уровня горизонта, Максимальный вылет стрелы, 21 м, Угол поворота люльки, 360 градусов вокруг оси',
    brand: 'КАМАЗ-43118',
    machine: 'АГП-45-5К',
    sizeTs: ['11.40', '2.50', '3.93'],
    width: '5.70',
    mass: '19.50',
    sizeCradle: ['1.60', '1.00'],
    special: ' Шасси полноприводное, проходимое, подходит для бездорожья.',
    rostechReg: true,
  }
]

export const taxCoef = 1.2
