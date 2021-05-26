import styled from "styled-components";

const HelloWorld = styled.div.attrs({ children: "Hello MY World!" })<{
  $someFlag?: boolean;
}>`
  display: flex;
  margin: auto;

  ${({ $someFlag }) => $someFlag && ``}
`;

const App = () => {
  return (
    <HelloWorld></HelloWorld>
  );
};

export default App;