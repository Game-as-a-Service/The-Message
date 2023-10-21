import { useRecoilState } from 'recoil';
import { useMemo } from 'react';
import { LineUpCardsState } from '@/states/globalCards/lineUpCardStates';

export const useSingleLineUpCardByName = (cardId:string) => {
  const [lineUpCards] = useRecoilState(LineUpCardsState);

  const findSingleCard = useMemo(() => {
    const targetCard = lineUpCards.CardInformation.find((card) => card.cardId === cardId);
    return targetCard;
  }, [cardId, lineUpCards.CardInformation]);

  return findSingleCard;
};
