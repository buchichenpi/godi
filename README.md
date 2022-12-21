This is a wrapper of facebookarchive/inject which is a tool of di used in Go.

这是一个封装了facebookarchive/inject的依赖注入项目

注意事项：
    依照facebookarchive/inject的要求，所有需要注入的字段必须对其他包可见

使用方式：
只需要两步：
    1. godiInject.Provide(objs...)
    2. godiInject.Populate()
其中步骤1可以在任何地方调用，但是建议在一个固定的初始化文件中执行。步骤一中的objs可以是普通的实例，
也可以实现Constructor.Init方法并在Init方法中进行一些必要的初始化，注意在Init中不要新初始化一个
实例。

具体的使用案例可以查看test目录。

