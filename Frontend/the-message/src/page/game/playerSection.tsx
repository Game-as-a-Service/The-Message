import { useRecoilState } from "recoil";
import lineUpCardBack from "@/assets/images/backOfLineUpCard.jpg";
import functionalCardBack from "@/assets/images/backOfFunctionCard.jpg";
import Card from "@/components/items/card";
import { useFunctionalCard } from "@/hooks/useFunctionCard";
import { useEffect } from "react";
import { PersonalFunctionalCards, PersonalFunctionalCardsState } from "@/states/personalCards/personalFunctionalCards";
import { usePersonalLineUpCard } from "@/states/personalCards/personalLineUpCard";


const playerSection = () => { 

  const [personalFunctionalCards] = useRecoilState(PersonalFunctionalCardsState);

  const personalFunctionalCardBackUrl = personalFunctionalCards.cardKind === "LineUp" ? lineUpCardBack : functionalCardBack;

  const { getPersonFunctionalCard } = useFunctionalCard();
  const { personalLineUpCard } = usePersonalLineUpCard();

  const cardWidth = 128;
  const gap = 30;
  const singleCardOffset = cardWidth + gap;
  const handCardLength = (PersonalFunctionalCards.length > 5) ? 5 : PersonalFunctionalCards.length ;
  const cardsSectionOffset = ((cardWidth + gap) * (handCardLength -1)) / 2;

  useEffect(() => {
      const cardList = [ ...PersonalFunctionalCards.slice(0,5) ];
      let count = 0
      const interval = setInterval(() => {
          const card = cardList.shift();
          if ( !card ) { clearInterval(interval); return }
          const cardObj = getPersonFunctionalCard(card.cardId);

          cardObj.move(cardsSectionOffset - singleCardOffset * count, -10);
          setTimeout(() => cardObj.flip(), 400);
          count++;
      }, 200);
  }, []);

  return (
    <>
      <div className="w-full h-1/4 flex bottom-0 absolute">
        <div className="w-1/5">
          <div>
            {personalFunctionalCards.CardInformation.map((card) => (
              <Card
                  key={card.cardId}
                  {...card}
                  cardFrontUrl={card.cardUrl}
                  cardBackUrl={personalFunctionalCardBackUrl}
                  additionCSS={`hover:-translate-y-20`}
              />
            ))}
          </div>
        </div>
        <div className="w-3/5"></div>
        <div className="w-1/5">
            <Card key={personalLineUpCard.cardId} {...personalLineUpCard} cardFrontUrl={personalLineUpCard.cardUrl} cardBackUrl={lineUpCardBack} />
        </div>
      </div>
    </>
  );

};

export default playerSection;
