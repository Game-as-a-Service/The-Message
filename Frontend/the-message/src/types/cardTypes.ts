export type CardInformation = {
    cardId: string,
    cardUrl: string,
    cardState: 'open' | 'close',
    cardCoordinate: CardCoordinate,
    canBeTurn: boolean
}

export type CardSet = {
    CardInformation: CardInformation[],
    cardKind: 'LineUp' | 'Functional'
}

export type CardCoordinate = {
    currentCoordinateX: number,
    currentCoordinateY: number
}
