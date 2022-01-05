const MOVES: [(i32, i32); 8] = [
    (-1, 2),
    (-2, 1),
    (-2, -1),
    (-1, -2),
    (1, -2),
    (2, -1),
    (2, 1),
    (1, 2),
];

fn next(x: usize, y: usize, n: i32, mut board: [[i32; 8]; 8]) -> bool {
    board[x][y] = n; // mark it moved

    if n == 64 {
        print_board(board);
        return true;
    }

    for m in MOVES.iter() {
        let (xd, yd) = m;
        let x1 = x as i32 + xd;
        let y1 = y as i32 + yd;

        if 0 <= x1 && x1 < 8 && 0 <= y1 && y1 < 8 {
            let x2 = x1 as usize;
            let y2 = y1 as usize;

            if board[x2][y2] == 0 {
                if next(x2, y2, n + 1, board) {
                    return true; // done here
                }

                board[x2][y2] = 0; // reverse it
            }
        }
    }
    false
}

fn print_board(board: [[i32; 8]; 8]) {
    for i in 0..8 {
        for j in 0..8 {
            print!("{:3}", board[i][j])
        }
        println!();
    }
}

fn main() {
    let board = [[0; 8]; 8];
    next(0, 0, 1, board);
}
