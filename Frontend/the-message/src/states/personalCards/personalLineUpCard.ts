import { atom, useRecoilState } from "recoil";
import { CardInformation, CardSet } from "@/types/cardTypes";
import lurking from "@/assets/images/lineUp/lurking.jpg";

export const PersonalLineUpCards: CardInformation[] = [
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: lurking,
      cardState: "open",
      cardCoordinate: { currentCoordinateX: 630, currentCoordinateY: 50 },
      canBeTurn: true,
    }
];

const PersonalLineUpCardsSet: CardSet = {
    CardInformation: PersonalLineUpCards,
    cardKind: "LineUp",
};

export const PersonalLineUpCardsState = atom({
    key: "PersonalLineUpCard",
    default: PersonalLineUpCardsSet,
});


export const usePersonalLineUpCard = () => {
    const [ personalLineUpCards ] = useRecoilState( PersonalLineUpCardsState );
    return { personalLineUpCard : personalLineUpCards.CardInformation[0] };
};
