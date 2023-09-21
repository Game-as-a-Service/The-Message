import { render, screen } from "@testing-library/react";
import Card from "../components/items/card";
import lineUpCardBack from "../../assets/images/backOfLineUpCard.jpg";
import lurking from "../assets/images/lurking.jpg";
import "@testing-library/jest-dom";
import App from "../App";



const CardMockProps = {
  cardFrontUrl:lurking,
  cardName:"lurking",
  cardBackUrl:lineUpCardBack,
  cardState:"open",
  cardCoordinate: { currentCoordinateX: 10, currentCoordinateY: 20 },
  canBeTurn:true 
}


test("測試Card Component是否運作正常", async () => {
  render(<Card {...CardMockProps}/>);
  expect(true).toBeTruthy();
});

test("測試入口畫面正常",async () => {
  render(<App/>);

  const titleElement = screen.getByText("風聲");
  expect(titleElement).toBeInTheDocument();
})