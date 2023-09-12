

interface CardProps {
  cardUrl:string,
  cardName:string,
  cardBackUrl:string,
}

const Card:React.FC<CardProps> = ( { cardUrl, cardBackUrl } ) => {

    return (
      <>
        <div className="flex min-h-screen items-center justify-center bg-slate-100">
          <div className="group h-96 w-80 [perspective:1000px]">
            <div className="relative h-full w-full rounded-xl shadow-xl transition-all duration-500 [transform-style:preserve-3d] group-hover:[transform:rotateY(540deg)]">
              <div className="absolute inset-0">
                <img className="card-image h-full w-full rounded-xl object-cover shadow-xl shadow-black/40" src={ cardUrl } style={{ border:"10px solid #000"}}/>
              </div>
              <div className="absolute inset-0 h-full w-full rounded-xl bg-black/80 px-12 text-center text-slate-200 [transform:rotateY(180deg)] [backface-visibility:hidden]">
                <div className="absolute inset-0">
                  <img className="card-image h-full w-full rounded-xl object-cover shadow-xl shadow-black/40" src={ cardBackUrl } style={{ border:"10px solid #000"}}/>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
    );
}

export default Card;