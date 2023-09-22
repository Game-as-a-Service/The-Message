import { useRecoilState } from "recoil";
import { FunctionalCardsState }  from "../states/functionalCardStates";
import { useMemo } from "react";

export const useSingleFunctionalCardByName = ( cardName:string ) => {

    const [ functionalCards ] = useRecoilState(FunctionalCardsState); 

    const findSingleCard = useMemo(() => {
        const targetCard = functionalCards.CardInformation.find(card => card.cardName === cardName);
        return targetCard;
    },[cardName, functionalCards.CardInformation]);

    return findSingleCard;
}