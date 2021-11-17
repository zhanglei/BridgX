package model

// ClusterTag
//使用Tags来描述Cluster的用途，属性等
// 比如使用
//	{
//		"group": "serviceName"
//  	"imageId": "xxx:1.2"
//  }
type ClusterTag struct {
	Base
	ClusterName string
	TagKey      string
	TagValue    string
}

func (ClusterTag) TableName() string {
	return "cluster_tag"
}

func GetTagsByClusterName(clusterName string) ([]ClusterTag, error) {
	clusterTags := make([]ClusterTag, 0)
	err := QueryAll(map[string]interface{}{"cluster_name": clusterName}, &clusterTags, "")
	if err != nil {
		logErr("GetTagsByClusterName from read db", err)
		return nil, err
	}
	return clusterTags, nil
}
