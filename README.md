<div id="top"></div>

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <img src="https://camo.qiitausercontent.com/83b135f2a0f5b083eccb132a35d3b40d72ebe34c/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d476f2d3736453146452e7376673f6c6f676f3d676f267374796c653d666f722d7468652d6261646765">
</p>

## 目次

1. [プロジェクトについて](#プロジェクトについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [開発環境構築](#開発環境構築)
5. [トラブルシューティング](#トラブルシューティング)

<!-- READMEの作成方法のドキュメントのリンク -->
<br />
<div align="right">
    <a href="READMEの作成方法のリンク"><strong>READMEの作成方法 »</strong></a>
</div>
<br />
<!-- Dockerfileのドキュメントのリンク -->
<div align="right">
    <a href="Dockerfileの詳細リンク"><strong>Dockerfileの詳細 »</strong></a>
</div>
<br />
<!-- プロジェクト名を記載 -->

## プロジェクト名

GET /para のサンプルコード

<!-- プロジェクトについて -->

## プロジェクトについて

パラメータ付き GET のサンプル

<!-- プロジェクトの概要を記載 -->

  <p align="left">
    <br />
    <!-- プロジェクト詳細にBacklogのWikiのリンク -->
    <a href="Backlogのwikiリンク"><strong>プロジェクト詳細 »</strong></a>
    <br />
    <br />

<p align="right">(<a href="#top">トップへ</a>)</p>

## 環境

<!-- 言語、フレームワーク、ミドルウェア、インフラの一覧とバージョンを記載 -->

| 言語・フレームワーク  | バージョン |
| --------------------- | ---------- |
| Go                | ?     |


<p align="right">(<a href="#top">トップへ</a>)</p>

## ディレクトリ構成

<!-- Treeコマンドを使ってディレクトリ構成を記載 -->

<pre>
.
├── README.md
├── handler
│   ├── go.mod
│   └── router.go
├── main
│   ├── go.mod
│   └── main.go
└── templete_README.md
</pre>

<p align="right">(<a href="#top">トップへ</a>)</p>

## 開発環境構築

...


### 動作確認

curl http://localhost:3000/para -d starttime=1701356400000 -d endtime=1704034799999 -d -d datasource[0]="V2H" -d datasource[1]="PV" -G

参考：[curlコマンド](https://ichiroku11.hatenablog.jp/entry/2014/03/31/232758)

参考：[クエリストリングで配列を表現をするケースをざっと調べる](https://shinkufencer.hateblo.jp/entry/2019/05/18/163912)

参考：[よく使うcurlコマンドのオプションまとめ（14個）](https://qiita.com/shtnkgm/items/45b4cd274fa813d29539)


<p align="right">(<a href="#top">トップへ</a>)</p>
