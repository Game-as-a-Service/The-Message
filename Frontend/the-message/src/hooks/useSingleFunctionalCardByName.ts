import { useRecoilState } from 'recoil';
import { useMemo } from 'react';
import { FunctionalCardsState } from '@/states/globalCards/functionalCardStates';

export const useSingleFunctionalCardByName = (cardId:string) => {
  const [functionalCards] = useRecoilState(FunctionalCardsState);

  const findSingleCard = useMemo(() => {
    const targetCard = functionalCards.CardInformation.find((card) => card.cardId === cardId);
    return targetCard;
  }, [cardId, functionalCards.CardInformation]);

  return findSingleCard;
};
