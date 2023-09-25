import { useRecoilState } from 'recoil';
import Card from '@/components/items/card';
import { LineUpCardsState } from '@/states/lineUpCardStates';
import lineUpCardBack from '@/assets/images/backOfLineUpCard.jpg';
import functionalCardBack from '@/assets/images/backOfFunctionCard.jpg';
import { FunctionalCardsState } from '@/states/functionalCardStates';

const Map = () => {
  const [lineUpCards] = useRecoilState(LineUpCardsState);
  const [functionalCards] = useRecoilState(FunctionalCardsState);

  const lineUpCardBackUrl = (lineUpCards.cardKind === 'LineUp') ? lineUpCardBack : functionalCardBack;
  const functionalCardBackUrl = (functionalCards.cardKind === 'Functional') ? functionalCardBack : lineUpCardBack;

  return (
    <div className="h-full w-full" style={{ position: 'relative' }}>
        {
            lineUpCards.CardInformation.map((card) => (
                <Card key={card.cardId} {...card} cardFrontUrl={card.cardUrl} cardBackUrl={lineUpCardBackUrl} />
            ))
        }
        {
            functionalCards.CardInformation.map((card) => (
                <Card key={card.cardId} {...card} cardFrontUrl={card.cardUrl} cardBackUrl={functionalCardBackUrl} />
            ))
        }
    </div>
  );
}

export default Map;
