import { atom } from 'recoil';
import { CardInformation, CardSet } from '@/types/cardTypes';
import millitaryIntelligence from '@/assets/images/lineUp/millitaryIntelligence.jpg';

const LineUpCards: CardInformation[] = [
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: millitaryIntelligence,
      cardState: 'close',
      cardCoordinate: { currentCoordinateX: 0, currentCoordinateY: 0 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: millitaryIntelligence,
      cardState: 'close',
      cardCoordinate: { currentCoordinateX: 5, currentCoordinateY: 0 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: millitaryIntelligence,
      cardState: 'close',
      cardCoordinate: { currentCoordinateX: 10, currentCoordinateY: 0 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: millitaryIntelligence,
      cardState: 'close',
      cardCoordinate: { currentCoordinateX: 15, currentCoordinateY: 0 },
      canBeTurn: false,
    },
];

const LineUpCardSet: CardSet = {
  CardInformation: LineUpCards,
  cardKind: 'LineUp',
};

export const LineUpCardsState = atom({
  key: 'lineUpCards',
  default: LineUpCardSet,
});
