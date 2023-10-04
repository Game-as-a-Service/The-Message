import { usePersonalFunctionalCards } from './../states/personalCards/personalFunctionalCards';
import type { CardInformation } from '@/types/cardTypes'

export const useFunctionalCard = () => {

    const { updateCardById, getFunctionalCardById } = usePersonalFunctionalCards();

    const getPersonFunctionalCard = (cardId: string) => {
        const card = getFunctionalCardById(cardId)
        return {
            ...card,
            move: (x: number, y: number) => {
                updateCardById(cardId, {
                    cardCoordinate: {
                      currentCoordinateX: x,
                      currentCoordinateY: y,
                    }
                });
            },

            flip: (cardState: CardInformation['cardState'] = card?.cardState === 'open' ? 'close' : 'open') => {
                updateCardById(cardId, {cardState})
            },

            putCardOnTheTable: () => {
                if(! card) return false;
                const tablePosition = { coordinateX: 10, coordinateY: 20 };
                updateCardById(cardId, {
                    cardCoordinate: {
                        currentCoordinateX: tablePosition.coordinateX,
                        currentCoordinateY: tablePosition.coordinateY, 
                    }
                })
            }
        }
    }

    return { getPersonFunctionalCard };

};
