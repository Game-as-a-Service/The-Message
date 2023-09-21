import { atom } from 'recoil';
import { CardInformation, CardSet } from './../types/cardTypes';
import millitaryIntelligence from "../assets/images/millitaryIntelligence.jpg";

const LineUpCards: CardInformation[] = [
    {
        cardId: "8b2ab4ca-615b-4479-a886-5d7f0624f853",
        cardUrl: millitaryIntelligence,
        cardName: "millitaryIntelligence",
        cardState: "close",
        cardCoordinate: { currentCoordinateX: 0 , currentCoordinateY: 0 },
        canBeTurn: true
    }
];


const LineUpCardSet: CardSet = {
    CardInformation: LineUpCards,
    cardKind: "LineUp"
}

export const LineUpCardsState = atom({
    key: "lineUpCards",
    default: LineUpCardSet,
})

