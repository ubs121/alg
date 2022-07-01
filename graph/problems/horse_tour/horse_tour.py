
class HorseTour:
    def __init__(self):
        self.board=[[0 for i in range(8)] for j in range(8)] # 8x8 board
        self.done=False

    # NOTE: not a full cycle
    def move(self, x:int, y:int, n:int):
        if not self.done and 0<=x<8 and 0<=y<8 and self.board[x][y]==0:
            self.board[x][y]=n # mark it visited
            if n==64:
                self.printBoard()
                self.done=True
                return
            self.move(x-1, y+2, n+1)
            self.move(x-2, y+1, n+1)
            self.move(x-2, y-1, n+1)
            self.move(x-1, y-2, n+1)
            self.move(x+1, y-2, n+1)
            self.move(x+2, y-1, n+1)
            self.move(x+2, y+1, n+1)
            self.move(x+1, y+2, n+1)
            
            self.board[x][y] = 0

    def printBoard(self):
        for i in range(8):
            for j in range(8):
                print("{}".format(self.board[i][j]).ljust(3), end = ' ')
            print("")

# testing 
h=HorseTour()
h.move(0,0,1)