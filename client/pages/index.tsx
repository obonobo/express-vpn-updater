import styled from "styled-components";
import Navbar from "../src/components/Navbar";
import HomeBar from "../src/components/Home";

const Content = styled.div`
  justify-content: center;
  flex-direction: column;
  justify-items: center;
  display: flex;

  & > * {
    margin-top: 7rem;
  }
`;

const Title = styled.text.attrs({ children: "EXPRESS VPN UPDATER" })`
  color: red;
  font-size: 280%;
  font-family: Roboto, sans-serif;
  padding: 1%;
`;



const Home = () => (
  <main>
    <Navbar>
      <Title />
    </Navbar>
    <div style={{height : "50px"}}/>
    <HomeBar/>
  </main>
);

export default Home;
