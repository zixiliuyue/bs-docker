## 国际化

前端国际化脚本 `vue-i18n-extract.js` 在每次运行`dev`模式时会自动检测文案变动，然后做`新增`和`删除`操作。

### 注意事项

1. 国际化文件中不能含有`特殊字符(引号，点等)`，不然脚本会自动添加`转义字符`，导致自动化失败
2. 过长文字请使用`xx提示语`作为key
3. 所有国际化文案`无需加规则前缀`(为了简化书写)，最终所有文件都会打平，放入对应目录即可，便于校对


### 文件规则

- unknow 未分类的文案
- button 按钮类场景
- content 界面上显示的文本内容
- important 触达用户的内容，需要精细化翻译
- label 表单名称、table表头、下拉选项、复选、单选、弹窗主体

### 翻译规范

