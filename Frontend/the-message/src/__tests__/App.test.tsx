import { render } from "@testing-library/react";
import Card from "@/components/items/card";
import lineUpCardBack from "@/assets/images/backOfLineUpCard.jpg";
import lurking from "@/assets/images/lurking.jpg";
import "@testing-library/jest-dom";


const CardMockProps = {
  cardFrontUrl:lurking,
  cardBackUrl:lineUpCardBack,
  cardState:"open",
  cardCoordinate: { currentCoordinateX: 10, currentCoordinateY: 20 },
  canBeTurn:true 
}


test("測試Card Component是否運作正常", async () => {
  render(<Card {...CardMockProps}/>);
  expect(true).toBeTruthy();
});

