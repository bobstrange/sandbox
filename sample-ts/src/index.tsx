import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

type SquareType = 'X' | 'O' | null
type SquareProps = {
  value: SquareType
  onClick: () => void
}

class Square extends React.Component<SquareProps> {
  render() {
    return (
      <button
        className="square"
        onClick={() => this.props.onClick()}
      >
        {this.props.value}
      </button>
    );
  }
}

type BoardProps = {
  // squares: SquareType[]
}
type BoardState = {
  squares: SquareType[]
}

class Board extends React.Component<BoardProps, BoardState> {
  constructor(props: BoardProps) {
    super(props)
    this.state = {
      squares: Array(9).fill(null)
    }
  }
  handleClick(i: number) {
    const squares = [...this.state.squares]
    squares[i] = 'X'
    this.setState({ squares })
  }

  renderSquare(i: number) {
    return <Square
      value={this.state.squares[i]}
      onClick={() => { this.handleClick(i) }}
    />;
  }

  render() {
    const status = 'Next player: X';

    return (
      <div>
        <div className="status">{status}</div>
        <div className="board-row">
          {this.renderSquare(0)}
          {this.renderSquare(1)}
          {this.renderSquare(2)}
        </div>
        <div className="board-row">
          {this.renderSquare(3)}
          {this.renderSquare(4)}
          {this.renderSquare(5)}
        </div>
        <div className="board-row">
          {this.renderSquare(6)}
          {this.renderSquare(7)}
          {this.renderSquare(8)}
        </div>
      </div>
    );
  }
}

class Game extends React.Component {
  render() {
    return (
      <div className="game">
        <div className="game-board">
          <Board />
        </div>
        <div className="game-info">
          <div>{/* status */}</div>
          <ol>{/* TODO */}</ol>
        </div>
      </div>
    );
  }
}

// ========================================

ReactDOM.render(
  <Game />,
  document.getElementById('root')
);
