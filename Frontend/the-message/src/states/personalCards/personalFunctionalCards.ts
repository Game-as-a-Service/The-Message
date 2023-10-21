import { atom, useRecoilState } from "recoil";
import { CardInformation, CardSet } from "@/types/cardTypes";
import  burnDown from '@/assets/images/functional/burnDown.jpg';
import misdirection from "@/assets/images/functional/misdirection.jpg";
import decode from "@/assets/images/functional/decode.jpg";
import intercept from "@/assets/images/functional/intercept.jpg";
import lockup from "@/assets/images/functional/lockup.jpg";

export const PersonalFunctionalCards: CardInformation[] = [
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: burnDown,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -580, currentCoordinateY: 50 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: misdirection,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -585, currentCoordinateY: 50 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: decode,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -590, currentCoordinateY: 50 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: intercept,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -595, currentCoordinateY: 50 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: lockup,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -600, currentCoordinateY: 50 },
      canBeTurn: false,
    },
    {
      cardId: window.crypto.randomUUID(),
      cardUrl: lockup,
      cardState: "close",
      cardCoordinate: { currentCoordinateX: -630, currentCoordinateY: 50 },
      canBeTurn: false,
    }
];

const PersonalFunctionalCardsSet: CardSet = {
    CardInformation: PersonalFunctionalCards,
    cardKind: "Functional",
};

export const PersonalFunctionalCardsState = atom({
    key: "PersonalLineUpCards",
    default: PersonalFunctionalCardsSet,
});


export const usePersonalFunctionalCards = () => {

  const [personalFunctionalCards, setPersonalLineUpCards] = useRecoilState( PersonalFunctionalCardsState );

  const updateCardById = (cardId: string, newCard: Partial<CardInformation>) => {
      setPersonalLineUpCards((prevCards)=>{
          const newCardsData = (() => {
            return {
                ...prevCards,
                CardInformation: prevCards.CardInformation.map((card) => {
                    if (card.cardId === cardId) { return { ...card, ...newCard } }
                    else { return { ...card } }
                }),
            };
          })();
          return newCardsData
      });
  }

  return {
      personalFunctionalCards,
      getFunctionalCardById: (cardId:string) => { return personalFunctionalCards.CardInformation.find(item => item.cardId === cardId)},
      updateCardById: updateCardById
  };
};
