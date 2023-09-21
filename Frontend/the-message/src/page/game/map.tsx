import Card from "../../components/items/card";
import { LineUpCardsState }  from "../../states/lineUpCardStates";
import { useRecoilState } from "recoil";
import lineUpCardBack from "../../assets/images/backOfLineUpCard.jpg";
import functionalCardBack from "../../assets/images/backOfFunctionCard.jpg";

const Map = () => {

    const [ lineUpCards ] = useRecoilState(LineUpCardsState); 
    
    return (
        <div className="h-full w-full" style={{position:'relative'}}>
            {
                lineUpCards.CardInformation.map((card) => (
                    <Card key={card.cardName} cardName={card.cardName} cardFrontUrl={card.cardUrl} canBeTurn={card.canBeTurn} cardState={card.cardState} cardCoordinate={card.cardCoordinate} cardBackUrl={(lineUpCards.cardKind == "LineUp") ? lineUpCardBack : functionalCardBack } />
                ))
            }
        </div>
    );

};

export default Map;

