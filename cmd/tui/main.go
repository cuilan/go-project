package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// ConfigOption 代表一个配置项
type ConfigOption struct {
	Text     string          // 显示的文本
	Enabled  bool            // 是否启用
	IsLeaf   bool            // 是否是叶子节点（可配置项）
	Children []*ConfigOption // 子选项
}

func main() {
	// 创建根节点数据
	rootData := &ConfigOption{
		Text: "Application Root Config",
		Children: []*ConfigOption{
			{
				Text: "App setup",
				Children: []*ConfigOption{
					{Text: "Select environment profile", IsLeaf: true, Enabled: true},
					{Text: "dev", IsLeaf: true},
					{Text: "test", IsLeaf: true},
					{Text: "prod", IsLeaf: true},
				},
			},
			{
				Text: "Device Drivers",
				Children: []*ConfigOption{
					{Text: "Serial device support", IsLeaf: true},
					{
						Text: "Block devices",
						Children: []*ConfigOption{
							{Text: "Normal floppy disk support", IsLeaf: true},
							{Text: "Loopback device support", IsLeaf: true, Enabled: true},
						},
					},
				},
			},
		},
	}

	app := tview.NewApplication()

	// 创建根 TreeView 节点
	rootNode := tview.NewTreeNode(rootData.Text)
	tree := tview.NewTreeView().
		SetRoot(rootNode).
		SetCurrentNode(rootNode)

	// 递归函数，用于将数据添加到 TreeView
	var add func(target *tview.TreeNode, data *ConfigOption)
	add = func(target *tview.TreeNode, data *ConfigOption) {
		for _, childData := range data.Children {
			// 根据状态设置节点文本
			nodeText := childData.Text
			if childData.IsLeaf {
				if childData.Enabled {
					nodeText = fmt.Sprintf("[*] %s", childData.Text)
				} else {
					nodeText = fmt.Sprintf("[ ] %s", childData.Text)
				}
			}

			childNode := tview.NewTreeNode(nodeText).
				SetReference(childData). // 将原始数据引用附加到节点上
				SetSelectable(true)

			target.AddChild(childNode)

			if len(childData.Children) > 0 {
				add(childNode, childData)
			}
		}
	}

	// 填充整个树
	add(rootNode, rootData)

	// 当用户在一个节点上按回车时触发
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		// 获取我们之前附加的数据引用
		reference := node.GetReference()
		if reference == nil {
			return // 根节点没有引用
		}

		option := reference.(*ConfigOption)

		// 只处理叶子节点（可配置项）
		if option.IsLeaf {
			// 切换启用状态
			option.Enabled = !option.Enabled

			// 更新节点显示的文本
			var newNodeText string
			if option.Enabled {
				newNodeText = fmt.Sprintf("[*] %s", option.Text)
			} else {
				newNodeText = fmt.Sprintf("[ ] %s", option.Text)
			}
			node.SetText(newNodeText)
		} else {
			// 如果不是叶子节点，则展开或折叠它
			node.SetExpanded(!node.IsExpanded())
		}
	})

	// 设置底部帮助信息
	helpText := tview.NewTextView().
		SetDynamicColors(true).
		SetText(" [yellow]Enter[white]: Toggle / Expand | [yellow]Arrows[white]: Navigate | [yellow]Ctrl+C[white]: Exit")

	// 创建整体布局
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tree, 0, 1, true).     // tree占据主要空间
		AddItem(helpText, 1, 0, false) // helpText占据1行

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
