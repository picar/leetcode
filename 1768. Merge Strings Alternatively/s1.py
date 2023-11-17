class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:

        l1 = len(word1)
        l2 = len(word2)
        ml = max(l1, l2)
        s = ""
        for i in range(ml):
            if i < l1:
                s += word1[i]
            if i < l2:
                s += word2[i]
        return s

if __name__ == "__main__":
    word1 = "hello"
    word2 = "world"
    print(Solution().mergeAlternately(word1, word2))
