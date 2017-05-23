### 特定のaws profileを使用する

```
aws configure --profile <your profile name>
```

```
provider "aws" {
    profile = "profile name"
}
```

### 実行時のグラフを表示する

```
terraform graph -type=plan | dot -Tpng > graph.png
```

![./graph.png "test"]
