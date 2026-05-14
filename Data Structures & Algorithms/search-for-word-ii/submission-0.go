type Trie struct {
    Children [26]*Trie
    Word     string // Store the word here to avoid passing a string in recursion
}

func addWord(word string, root *Trie) {
    curr := root
    for _, char := range word {
        index := char - 'a'
        if curr.Children[index] == nil {
            curr.Children[index] = &Trie{}
        }
        curr = curr.Children[index]
    }
    curr.Word = word
}

func findWords(board [][]byte, words []string) []string {
    root := &Trie{}
    for _, w := range words {
        addWord(w, root)
    }

    var res []string
    for r := 0; r < len(board); r++ {
        for c := 0; c < len(board[0]); c++ {
            backtrack(r, c, board, root, &res)
        }
    }
    return res
}

func backtrack(r, c int, board [][]byte, node *Trie, res *[]string) {
    // 1. Bounds check and character match check
    if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
        return
    }
    
    char := board[r][c]
    if char == '#' || node.Children[char-'a'] == nil {
        return
    }

    // 2. Move to the child node
    node = node.Children[char-'a']
    
    // 3. Check if we found a word
    if node.Word != "" {
        *res = append(*res, node.Word)
        node.Word = "" // Avoid duplicates
    }

    // 4. Mark as visited using the board itself (saves memory vs a map)
    board[r][c] = '#'

    // 5. Explore neighbors
    backtrack(r+1, c, board, node, res)
    backtrack(r-1, c, board, node, res)
    backtrack(r, c+1, board, node, res)
    backtrack(r, c-1, board, node, res)

    // 6. Restore character (backtrack)
    board[r][c] = char
}