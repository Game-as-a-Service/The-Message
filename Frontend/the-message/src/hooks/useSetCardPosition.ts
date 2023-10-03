import { useRecoilState } from 'recoil';
import { CardInformation, CardCoordinate } from '../types/cardTypes';
import { LineUpCardsState } from '../states/lineUpCardStates';

export const useSetCardPosition = (setCoordinateX: number, setCoordinateY: number, cardId: string) => {
  const [lineUpCards, setCardContent] = useRecoilState(LineUpCardsState);

  const targetCard : CardInformation | undefined = lineUpCards.CardInformation.find((card) => card.cardId === cardId);
  if (!targetCard) return;

  targetCard.cardCoordinate = { currentCoordinateX: setCoordinateX, currentCoordinateY: setCoordinateY } as CardCoordinate;

  setCardContent(lineUpCards);
};
