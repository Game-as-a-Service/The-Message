import { useRecoilState } from "recoil";
import { LineUpCardsState }  from "../states/lineUpCardStates.ts";
import { useMemo } from "react";

export const useSingleLineUpCardByName = ( cardName:string ) => {

    const [ lineUpCards ] = useRecoilState(LineUpCardsState); 

    const findSingleCard = useMemo(() => {
        const targetCard = lineUpCards.CardInformation.find(card => card.cardName === cardName);
        return targetCard;
    },[cardName, lineUpCards.CardInformation]);

    return findSingleCard;
}