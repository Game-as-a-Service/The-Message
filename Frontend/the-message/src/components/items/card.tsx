import { CardCoordinate } from "../../types/cardTypes";
import clsx from "clsx";

interface CardProps {
  cardFrontUrl:string,
  cardName:string,
  cardBackUrl:string,
  cardState:string,
  cardCoordinate: CardCoordinate
  canBeTurn:boolean;
}

const Card:React.FC<CardProps> = ( { cardFrontUrl, cardBackUrl, cardState, canBeTurn, cardCoordinate } ) => {

    const cardPositionStyle = {
        top:`calc(50% + ${cardCoordinate.currentCoordinateX}px)`,
        left:`calc(50% + ${cardCoordinate.currentCoordinateY}px)`,
        transition:"all 0.5s ease-out",
        transform:"translate(-50%,-50%)"
    }
    
    const cardOnTheFront = (cardState === "open") ? cardFrontUrl : cardBackUrl ; 
    const cardOnTheBack = (cardState === "close") ? cardFrontUrl : cardBackUrl ; 

    return (
      <>
        <div className="absolute" style={ cardPositionStyle }>
          <div className="group h-80 w-64 [perspective:1000px]">
            <div className={clsx(`relative h-full w-full rounded-xl shadow-xl transition-all duration-500 [transform-style:preserve-3d]`, canBeTurn && "group-hover:[transform:rotateY(540deg)]")}>
              <div className="absolute inset-0">
                <img className="card-image h-full w-full rounded-xl object-cover shadow-xl shadow-black/40" src={ cardOnTheFront } style={{ border:"10px solid #000"}}/>
              </div>
              <div className="absolute inset-0 h-full w-full rounded-xl bg-black/80 px-12 text-center text-slate-200 [transform:rotateY(180deg)] [backface-visibility:hidden]">
                <div className="absolute inset-0">
                  <img className="card-image h-full w-full rounded-xl object-cover shadow-xl shadow-black/40" src={ cardOnTheBack } style={{ border:"10px solid #000"}}/>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
    );
}

export default Card;