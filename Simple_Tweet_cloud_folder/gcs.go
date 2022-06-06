package main

import (
    "context"
    "fmt"
    "io"

    "cloud.google.com/go/storage"
)

const (
    BUCKET_NAME = "around_gcs_bucket"
)

func saveToGCS(r io.Reader, objectName string)(string, error) {
    ctx := context.Background()
    //在相同账号中运行不同功能可以共用，如果是不同机器呢那就得加WithCredential了
    //我需要下载my own Credential来试试
    client, err := storage.NewClient(ctx)

    if err != nil {
        return "", err
    }

    //这个objectName一般来说是post的id, 
    object:= client.Bucket(BUCKET_NAME).Object(objectName)
    //开始写入storage
    wc := object.NewWriter(ctx)
    if _,err := io.Copy(wc, r); err != nil {
        return "", err
    }

    if err := wc.Close(); err != nil {
        return "", err
    }

    //设定bucket里面的ACLHandle在&bucketHandler里面的参数， 然后利用下面那个function进行权限设置
    //func (a *ACLHandle) Set(ctx context.Context, entity ACLEntity, role ACLRole)(err error)： sets the role for the given entity and return err.

    //ACLEntity: user or group.
    //ACLRole: level of access to grant.
    
    //storage.RoleReader is a string called "READER"
    //Access Control List--> user_group: All users, user_access: reader(read only)

    if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return "", err
    }

    attrs, err := object.Attrs(ctx)

    if err != nil {
        return "", err
    }
    
    fmt.Printf("Image is saved to GCS: %s\n", attrs.MediaLink)
    return attrs.MediaLink, nil

}

//https://cloud.google.com/storage/docs/deleting-objects#storage-delete-object-go
func deleteFromGCS(objectName string) error {
    ctx := context.Background()
    //在相同账号中运行不同功能可以共用，如果是不同机器呢那就得加WithCredential了
    //我需要下载my own Credential来试试
    client, err := storage.NewClient(ctx)

    if err != nil {
        return err
    }

    //objectname: post.Id
    object := client.Bucket(BUCKET_NAME).Object(objectName)

    _,err = object.Attrs(ctx)
    if err != nil {
        return err
    }

    err = object.Delete(ctx)

    return err
}

/*
func (b *BucketHandle) Object(name string) *ObjectHandle {
	retry := b.retry.clone()
	return &ObjectHandle{
		c:      b.c,
		bucket: b.name,
		object: name,
		acl: ACLHandle{
			c:           b.c,
			bucket:      b.name,
			object:      name,
			userProject: b.userProject,
			retry:       retry,
		},
		gen:         -1,
		userProject: b.userProject,
		retry:       retry,
	}
}
*/
