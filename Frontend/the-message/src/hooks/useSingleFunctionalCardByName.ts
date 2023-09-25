import { useRecoilState } from 'recoil';
import { useMemo } from 'react';
import { FunctionalCardsState } from '../states/functionalCardStates';

export const useSingleFunctionalCardByName = (cardName:string) => {
  const [functionalCards] = useRecoilState(FunctionalCardsState);

  const findSingleCard = useMemo(() => {
    const targetCard = functionalCards.CardInformation.find((card) => card.cardName === cardName);
    return targetCard;
  }, [cardName, functionalCards.CardInformation]);

  return findSingleCard;
};
