package optimization

func OptimizeFrontBack(dictionary []string) []string {
    keys := make(map[string]bool);
    slice := []string{};
 
    for _, password := range dictionary {
        if _, has := keys[password]; !has {
            keys[password] = true;
            slice = append(slice, password);
        }
    }

    return slice;
}