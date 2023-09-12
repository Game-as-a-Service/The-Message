// import Card from "../../items/card";
import lurking from "../../../assets/images/lurking1.jpg";
import soysauce from "../../../assets/images/soysauce.jpg";
import millitaryIntelligence from "../../../assets/images/millitaryIntelligence.jpg";
import useCardRenderer from "../../../hooks/useCardRenderer.tsx";
import { useState } from "react";

type CardInformations = {
    cardUrl:string,
    cardName:string
}

type CardSet = {
    cardInformations: CardInformations[],
    cardKind: "LineUp" | "Functional"
} 

const Map = () => {
    
    const LineUpCards: CardInformations[] = [
        {
            cardUrl: lurking,
            cardName: "潛伏戰線",
        },
        {
            cardUrl: soysauce,
            cardName: "打醬油"
        },
        {
            cardUrl: millitaryIntelligence,
            cardName: "軍情處"
        }
    ];

    const LineUpCardSet: CardSet = {
        cardInformations: LineUpCards,
        cardKind: "LineUp"
    }
    

    const [cardsSet, setCardsSet ] = useState<CardSet>(LineUpCardSet);
    

    return (
        <div className="flex">
            { useCardRenderer(cardsSet) }
        </div>
    );

};

export default Map;

