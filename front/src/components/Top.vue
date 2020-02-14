<template>
  <div class='top'>
    <div class='top-area'>
      <h1>echo</h1>
      <h2>完全匿名SNS</h2>
      <button @click="signIn">ログイン</button>
      <button @click="signUp">新規登録</button>
    </div>

    <div class='information-area'>
      <h2>これなに</h2>

      <h3>ここではあなたは何者でもありません</h3>
      <p>
        あなたの発言をあなたがしたものであるということを他のユーザーが知る方法はありません。<br/>
        また同時にここでは、あなたの年齢や性別、リアルでの立場、ネットでの知名度などが発言に付加されることはありません。あなたの発言はどこかの誰かが零したこだまのひとつにすぎません。
      </p>

      <h3>直接的なコミュニケーションをする場所ではありません</h3>
      <p>リプライや拡散といった機能は存在しません。引用はありますが他人から引用されてもあなたに通知が届くことはありません。</p>

      <h3>承認欲求は満たされないかもしれません</h3>
      <p>あなたの発言に付けられたお気に入り数はあなたにだけ表示されます。</p>

      <h3>禁止事項</h3>
      <ul>
        <li>自他を問わず発言者が何者かを明かす行為</li>
        <ul>
          <li>他の発言者が何者か分かったとしてもそれについて言及しないでください。</li>
          <li>自分が何者であるかを表明したり、投稿が特定個人によるものと判別できるようにはしないでください。</li>
          <li>ただし、Twitterなど他の場所で同一内容の発言をすることは禁止していません。</li>
        </ul>
        <li>犯罪またはそれを助長する行為</li>
        <li>サービス内外の他人に対する過度の誹謗中傷</li>
        <li>スパム行為 (bot投稿を含む)</li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'Top',
  methods: {
    signIn: async function () {
      this.$router.push('/signin')
    },
    signUp: async function () {
      this.$router.push('/signup')
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
.top-area {
  align-content: center;
}
.information-area {
  float: left;
  text-align: left;
}
</style>
