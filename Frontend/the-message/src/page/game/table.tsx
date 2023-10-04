import { useRecoilState } from 'recoil';
import Card from '@/components/items/card';
import { LineUpCardsState } from '@/states/globalCards/lineUpCardStates';
import lineUpCardBack from '@/assets/images/backOfLineUpCard.jpg';
import functionalCardBack from '@/assets/images/backOfFunctionCard.jpg';
// import { FunctionalCardsState } from '@/states/functionalCardStates';

const Table = () => {

    const [lineUpCards] = useRecoilState(LineUpCardsState);
    const lineUpCardBackUrl = (lineUpCards.cardKind === 'LineUp') ? lineUpCardBack : functionalCardBack;

    return (
        <div className="w-5/6 bg-slate-700 absolute left-1/2 -translate-x-1/2 m-0 rounded-full translate-y-6" style={{height:"75%"}}>
            {
                lineUpCards.CardInformation.map((card) => (
                    <Card key={card.cardId} {...card} cardFrontUrl={card.cardUrl} cardBackUrl={lineUpCardBackUrl} />
                ))
            }
        </div>
    );
}

export default Table;