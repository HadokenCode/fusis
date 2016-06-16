package storm

import (
	"github.com/boltdb/bolt"
	"github.com/fatih/structs"
)

// Init creates the indexes and buckets for a given structure
func (n *Node) Init(data interface{}) error {
	if !structs.IsStruct(data) {
		return ErrBadType
	}

	info, err := extract(data)
	if err != nil {
		return err
	}

	err = n.s.Bolt.Update(func(tx *bolt.Tx) error {
		bucket, err := n.CreateBucketIfNotExists(tx, info.Name)
		if err != nil {
			return err
		}

		for fieldName, idxInfo := range info.Indexes {
			switch idxInfo.Type {
			case tagUniqueIdx:
				_, err = NewUniqueIndex(bucket, []byte(indexPrefix+fieldName))
			case tagIdx:
				_, err = NewListIndex(bucket, []byte(indexPrefix+fieldName))
			default:
				err = ErrBadIndexType
			}

			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// Init creates the indexes and buckets for a given structure
func (s *DB) Init(data interface{}) error {
	return s.root.Init(data)
}
