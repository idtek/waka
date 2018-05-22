package generator

const codeTemplate = `// Generated by github.com/'zhouyukun0622'/waka/protoc/protoc-gen-cellnet
// Source: {{.FileName}}
// DO NOT EDIT!!!

using CellnetSDK.Meta;

namespace {{.Namespace}}
{
    /// <summary>
    /// {{.FileName}} 元数据提供者
    /// </summary>
    public class {{.MetaProviderClassName}}
    {
        /// <summary>
        /// 注册消息
        /// </summary>
        static public void RegisterAll()
        {
            {{range .MessageType}}
            MetaTable.RegisterMessageMeta("{{.FullName}}", {{.ID}}, new {{.FullName}}().GetType(), (d) => {{.FullName}}.Parser.ParseFrom(d));
            {{end}}
        }
    }
}
`
