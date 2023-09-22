import Card from "@/components/items/card";
import { LineUpCardsState }  from "@/states/lineUpCardStates";
import { useRecoilState } from "recoil";
import lineUpCardBack from "@/assets/images/backOfLineUpCard.jpg";
import functionalCardBack from "@/assets/images/backOfFunctionCard.jpg";

const Map = () => {

    const [ lineUpCards ] = useRecoilState(LineUpCardsState); 
    const cardBackUrl =(lineUpCards.cardKind === "LineUp") ? lineUpCardBack : functionalCardBack ;

    return (
        <div className="h-full w-full" style={{position:'relative'}}>
            {
                lineUpCards.CardInformation.map((card) => (
                    <Card key={card.cardId} {...card} cardFrontUrl={card.cardUrl} cardBackUrl={ cardBackUrl } />
                ))
            }
        </div>
    );

};

export default Map;

