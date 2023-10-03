import { atom } from 'recoil';
import { CardInformation, CardSet } from '@/types/cardTypes';
import lurking from '@/assets/images/lurking.jpg';

const FunctionalCards : CardInformation[] = [
  {
    cardId: '8b2ab4ca-615b-4479-a886-5d7f0624f553',
    cardUrl: lurking,
    cardState: 'close',
    cardCoordinate: { currentCoordinateX: 200, currentCoordinateY: 20 },
    canBeTurn: true,
  },
];

const FunctionalCardSet: CardSet = {
  CardInformation: FunctionalCards,
  cardKind: 'Functional',
};

export const FunctionalCardsState = atom({
  key: 'functionalCards',
  default: FunctionalCardSet,
});
