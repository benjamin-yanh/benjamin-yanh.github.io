# 对话交接文档 — SGLang 学习笔记项目

> 用途：在新对话里贴入本文件内容（或让 Claude 读取本文件），即可延续此项目的上下文。
> 更新时间：2026-07-02

---

## 1. 项目目标

帮助用户（入门层次：会用 LLM，但不熟 Transformer/注意力内部机制）**研究学习论文 SGLang**，产出一份带原文图示、图解动画、可交互图表的中文学习笔记。

- 论文：*SGLang: Efficient Execution of Structured Language Model Programs*（arXiv **2312.07104v2**, 2024-06）。
- 论文 PDF 位置：`assets/academic/2312.07104v2.pdf`（在工作文件夹内）。
- 论文三大核心：**RadixAttention**（基数树+LRU 复用 KV Cache）、**压缩有限状态机**（约束解码一步多 token）、**API 推测执行**；前端是嵌入 Python 的 DSL（gen/select/fork/join/image 等）。

## 2. 工作文件夹

`/Users/benjamin/Documents/github/benjamin-yanh.github.io/assets/academic/`（这是一个 github.io 站点的 assets 目录）。

## 3. 已交付文件（持久，均在上面的文件夹）

- **`SGLang_学习笔记.md`** — 主学习笔记（Markdown 版）。含：前置知识地图（第0节，分"必须懂/最好懂/了解即可"）、逐节讲解（1~8节：总览、编程模型、RadixAttention、压缩FSM、API推测执行、实验、相关工作、局限）、速记卡、10道自测题。**内嵌原文图片**（`img/fig1,2,3,4,5,9.png`）和 **Mermaid 图**（§0.3 Trie/Radix 对比、§0.4 依赖关系图）。含多个"专栏"块（self-consistency 详解、代码与模型如何配合、Trie vs Radix 澄清等）。
- **`SGLang_学习笔记.html`** — 主笔记的 HTML 版（**当前主要交付物**，用户主要看这个）。由 Markdown 经 pandoc 转换 + 后处理生成。**图 2/3/9 用的是用户自己提供的代码（已内联，无 iframe）**；图 1/4/5 及两张概念图是 Claude 画的 SVG。页宽限 900px，图 2/3 突破列宽加宽。
- **`SGLang_图解动画.html`** — 独立的交互动画页（6 个动画：自回归&KV Cache、前缀复用、RadixAttention 基数树分步、压缩FSM对比、API推测执行、自洽投票）。
- **`fig2_essay_judge.html`** — 用户提供的图2独立版（代码+黄色箭头标注，Tailwind+JS 动态画箭头）。
- **`fig3_radixattention.html`** — 用户提供的图3独立版（1400×980 canvas，九步基数树，用户已多次微调坐标，**这是最新版**）。
- **`fig9_sharing.html`** — 用户提供的图9独立版（含 SVG 图 + 导出PNG/SVG/对照按钮，Tailwind）。其"对照原图"引用 `image_615f17.png`（用户需自行放入才显示）。
- **`img/`** — 只保留 6 张原文裁图 `fig1.png fig2.png fig3.png fig4.png fig5.png fig9.png` + 3 张对比图 `对比_图2/3/9_原图vs重绘.png`。

## 4. 关键：HTML 的生成方式（重要！）

`SGLang_学习笔记.html` 不是手写的，是用一套**临时脚本在 VM /tmp 下生成**的：
- `/tmp/build3.py`：读 `SGLang_学习笔记.md` → pandoc 转 HTML → 把 Mermaid 代码块替换成 SVG、把 `<img src="img/figN.png">` 替换成 `figs.json` 里对应的图（SVG 或用户内联代码）→ 套用 `/tmp/tpl.html` 模板输出。
- `/tmp/figs.json`：存放每张图的最终 HTML/SVG 片段（fig1,fig4,fig5,fig9,fig_dep,fig_trie 为 Claude 的 SVG；fig2,fig3 为用户内联代码；fig9 现为用户 SVG）。
- `/tmp/tpl.html`：HTML 模板（含 CSS）。
- `/tmp/inline2.py`：从 `fig2_essay_judge.html`/`fig3_radixattention.html` 抽取代码，重建 figs.json 里的 fig2/fig3 内联块。

**⚠️ 这些 /tmp 脚本在新会话中会丢失（VM 重置）。** 所以在新对话里若要改 HTML：
- 简单改动 → **直接编辑 `SGLang_学习笔记.html`**。
- 或重建整套 pipeline（费时）。
- 修改正文内容 → 改 `SGLang_学习笔记.md` 后需重新走 pipeline 才能反映到 HTML（pipeline 已不在，需重建或手动同步）。

## 5. 各图当前状态（HTML 版）

| 图 | 内容 | 当前实现 |
|---|---|---|
| 图1 | 系统架构 | Claude 的 SVG |
| 图2 | 多维度作文评审代码+标注 | **用户代码内联**（`#fig2wrap`，Tailwind class 已改写为局部内联样式，左对齐，加宽突破列宽）|
| 图3 | 九步 RadixAttention 基数树 | **用户 canvas 内联**（`#myCanvas` 1400×980，加宽突破列宽）|
| 图4 | 普通 vs 压缩 FSM | Claude 的 SVG |
| 图5 | 吞吐柱状图 | Claude 的 SVG（柱高按原图近似读数）|
| 图9 | KV 共享四模式 | **用户 SVG 内联**（`#diagram-svg`，含 tree-of-thought 深层分支，自适应宽度）|
| §0.3 | Trie vs Radix 对比 | Claude 的 SVG |
| §0.4 | 依赖关系图 | Claude 的 SVG |

页宽：`.wrap{max-width:900px}`；图2/3 用 `figure:has(#myCanvas),figure:has(#code-container)` 加宽到 `min(1400px,94vw)` 居中、去边框。图9 目前随列宽自适应（未单独加宽）。

## 6. 用户偏好（务必遵守）

- **简洁直接**，少解释少赘述（用户在偏好里设定）。
- **不准反驳，改为询问**（用户明确要求）：遇到分歧/自己有更好方案时，不要直接反驳用户，先用提问确认。此前因反驳被批评过。
- 倾向**用户自己提供的代码/原图**，而非 Claude 的重绘；Claude 重绘只是"高度近似"。
- 图表**内联**进 HTML，不要用 iframe（用户明确要求过）。
- 中文回复。

## 7. 可能的后续任务（用户提过或自然延伸）

- 图 9 也像图 3 那样加宽（已提议，用户未明确要）。
- 图 1/4/5 是否也换成用户版本。
- 正文宽度/图注/样式微调。
- 用户常见模式：**发来一份 HTML/canvas/SVG 代码，要求"替换掉你画的图X"** → 做法：存为独立文件 + 抽取核心图形内联进笔记 HTML（去掉整页外壳/按钮/深色背景），保持无 iframe。

## 8. 环境备注

- 有 pandoc、ImageMagick(convert)、pdftoppm/pdftocairo/pdfimages 可用（VM 内）。
- VM 路径映射：工作文件夹 = `/sessions/<id>/mnt/academic/`；file 工具用 host 路径 `/Users/benjamin/.../assets/academic/`。
- 光栅化预览 SVG 时缺中文字体（中文显示空白），但浏览器正常——验证版式可用 convert，验证中文需靠浏览器。
- 删除文件需先调 `allow_cowork_file_delete`（曾遇到 rm "Operation not permitted"）。
