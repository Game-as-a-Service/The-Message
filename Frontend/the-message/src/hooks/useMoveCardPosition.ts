import { useRecoilState } from 'recoil';
import { CardInformation, CardCoordinate } from '../types/cardTypes';
import { LineUpCardsState } from '../states/lineUpCardStates';

export const useMoveCardPosition = (movementOnX: number, movementOnY: number, cardId: string) => {
  const [lineUpCards, setCardContent] = useRecoilState(LineUpCardsState);

  const targetCard : CardInformation | undefined = lineUpCards.CardInformation.find((card) => card.cardId === cardId);
  if (!targetCard) return;

  const { currentCoordinateX, currentCoordinateY } : CardCoordinate = targetCard.cardCoordinate;
  targetCard.cardCoordinate = { currentCoordinateX: currentCoordinateX + movementOnX, currentCoordinateY: currentCoordinateY + movementOnY } as CardCoordinate;

  setCardContent(lineUpCards);
};
