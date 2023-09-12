import Card from '../components/items/card';
import lineUpCardBack from "../assets/images/backOfLineUpCard.jpg";
import functionalCardBack from "../assets/images/backOfFunctionCard.jpg";

type CardInformations = {
    cardUrl:string,
    cardName:string
}

type CardSet = {
    cardInformations: CardInformations[],
    cardKind: string
} 

const useCardRenderer = (cardSet: CardSet) => {
  return (
    <div className="flex">
        {cardSet.cardInformations.map((card: CardInformations) => (
            <Card cardName={card.cardName} cardUrl={card.cardUrl} cardBackUrl={(cardSet.cardKind == "LineUp") ? lineUpCardBack : functionalCardBack } />
        ))}
    </div>
  );
};

export default useCardRenderer;