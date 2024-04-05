// Code generated by fastssz. DO NOT EDIT.
// Hash: 0efc9050075e3e3ec95caec1378e27ec38c9bc0f847b8d0afc4d72a1ae128f80
// Version: 0.1.3
package types

import (
	enginetypes "github.com/berachain/beacon-kit/mod/execution/types"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/kzg"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockDeneb object to a target array
func (b *BeaconBlockDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(84)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ProposerIndex))

	// Field (2) 'ParentBlockRoot'
	dst = append(dst, b.ParentBlockRoot[:]...)

	// Field (3) 'StateRoot'
	dst = append(dst, b.StateRoot[:]...)

	// Offset (4) 'Body'
	dst = ssz.WriteOffset(dst, offset)

	// Field (4) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 84 {
		return ssz.ErrSize
	}

	tail := buf
	var o4 uint64

	// Field (0) 'Slot'
	b.Slot = primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = primitives.ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'ParentBlockRoot'
	copy(b.ParentBlockRoot[:], buf[16:48])

	// Field (3) 'StateRoot'
	copy(b.StateRoot[:], buf[48:80])

	// Offset (4) 'Body'
	if o4 = ssz.ReadOffset(buf[80:84]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 84 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (4) 'Body'
	{
		buf = tail[o4:]
		if b.Body == nil {
			b.Body = new(BeaconBlockBodyDeneb)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) SizeSSZ() (size int) {
	size = 84

	// Field (4) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockDeneb object with a hasher
func (b *BeaconBlockDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	hh.PutUint64(uint64(b.ProposerIndex))

	// Field (2) 'ParentBlockRoot'
	hh.PutBytes(b.ParentBlockRoot[:])

	// Field (3) 'StateRoot'
	hh.PutBytes(b.StateRoot[:])

	// Field (4) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockBodyDeneb object to a target array
func (b *BeaconBlockBodyDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Field (0) 'RandaoReveal'
	dst = append(dst, b.RandaoReveal[:]...)

	// Field (1) 'Graffiti'
	dst = append(dst, b.Graffiti[:]...)

	// Offset (2) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Deposits); ii++ {
		offset += 4
		offset += b.Deposits[ii].SizeSSZ()
	}

	// Offset (3) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Offset (4) 'BlobKzgCommitments'
	dst = ssz.WriteOffset(dst, offset)

	// Field (2) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.Deposits", size, 16)
		return
	}
	{
		offset = 4 * len(b.Deposits)
		for ii := 0; ii < len(b.Deposits); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Deposits[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	if size := len(b.BlobKzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.BlobKzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.BlobKzgCommitments); ii++ {
		dst = append(dst, b.BlobKzgCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o4 uint64

	// Field (0) 'RandaoReveal'
	copy(b.RandaoReveal[:], buf[0:96])

	// Field (1) 'Graffiti'
	copy(b.Graffiti[:], buf[96:128])

	// Offset (2) 'Deposits'
	if o2 = ssz.ReadOffset(buf[128:132]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'ExecutionPayload'
	if o3 = ssz.ReadOffset(buf[132:136]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'BlobKzgCommitments'
	if o4 = ssz.ReadOffset(buf[136:140]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (2) 'Deposits'
	{
		buf = tail[o2:o3]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Deposits[indx] == nil {
				b.Deposits[indx] = new(Deposit)
			}
			if err = b.Deposits[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'ExecutionPayload'
	{
		buf = tail[o3:o4]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (4) 'BlobKzgCommitments'
	{
		buf = tail[o4:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.BlobKzgCommitments = make([]kzg.Commitment, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKzgCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) SizeSSZ() (size int) {
	size = 140

	// Field (2) 'Deposits'
	for ii := 0; ii < len(b.Deposits); ii++ {
		size += 4
		size += b.Deposits[ii].SizeSSZ()
	}

	// Field (3) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (4) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockBodyDeneb object with a hasher
func (b *BeaconBlockBodyDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(b.RandaoReveal[:])

	// Field (1) 'Graffiti'
	hh.PutBytes(b.Graffiti[:])

	// Field (2) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	// Field (3) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	{
		if size := len(b.BlobKzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.BlobKzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKzgCommitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.BlobKzgCommitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(56)

	// Offset (0) 'Pubkey'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.Pubkey)

	// Field (1) 'Credentials'
	dst = append(dst, d.Credentials[:]...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(d.Amount))

	// Offset (3) 'Signature'
	dst = ssz.WriteOffset(dst, offset)

	// Field (4) 'Index'
	dst = ssz.MarshalUint64(dst, d.Index)

	// Field (0) 'Pubkey'
	dst = append(dst, d.Pubkey[:]...)

	// Field (3) 'Signature'
	dst = append(dst, d.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 56 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o3 uint64

	// Offset (0) 'Pubkey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 56 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Credentials'
	copy(d.Credentials[:], buf[4:36])

	// Field (2) 'Amount'
	d.Amount = primitives.Gwei(ssz.UnmarshallUint64(buf[36:44]))

	// Offset (3) 'Signature'
	if o3 = ssz.ReadOffset(buf[44:48]); o3 > size || o0 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'Index'
	d.Index = ssz.UnmarshallUint64(buf[48:56])

	// Field (0) 'Pubkey'
	{
		buf = tail[o0:o3]
		copy(d.Pubkey[:], buf)
	}

	// Field (3) 'Signature'
	{
		buf = tail[o3:]
		copy(d.Signature[:], buf)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 56

	// Field (0) 'Pubkey'
	size += len(d.Pubkey)

	// Field (3) 'Signature'
	size += len(d.Signature)

	return
}

// HashTreeRoot ssz hashes the Deposit object
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit object with a hasher
func (d *Deposit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Pubkey[:]))
		if byteLen > 48 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(d.Pubkey[:])
		hh.MerkleizeWithMixin(elemIndx, byteLen, (48+31)/32)
	}

	// Field (1) 'Credentials'
	hh.PutBytes(d.Credentials[:])

	// Field (2) 'Amount'
	hh.PutUint64(uint64(d.Amount))

	// Field (3) 'Signature'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Signature[:]))
		if byteLen > 96 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(d.Signature[:])
		hh.MerkleizeWithMixin(elemIndx, byteLen, (96+31)/32)
	}

	// Field (4) 'Index'
	hh.PutUint64(d.Index)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Deposit object
func (d *Deposit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}

// MarshalSSZ ssz marshals the DepositMessage object
func (d *DepositMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the DepositMessage object to a target array
func (d *DepositMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(44)

	// Offset (0) 'Pubkey'
	dst = ssz.WriteOffset(dst, offset)

	// Field (1) 'Credentials'
	dst = append(dst, d.Credentials[:]...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(d.Amount))

	// Field (0) 'Pubkey'
	dst = append(dst, d.Pubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the DepositMessage object
func (d *DepositMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 44 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Pubkey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 44 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Credentials'
	copy(d.Credentials[:], buf[4:36])

	// Field (2) 'Amount'
	d.Amount = primitives.Gwei(ssz.UnmarshallUint64(buf[36:44]))

	// Field (0) 'Pubkey'
	{
		buf = tail[o0:]
		copy(d.Pubkey[:], buf)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the DepositMessage object
func (d *DepositMessage) SizeSSZ() (size int) {
	size = 44

	// Field (0) 'Pubkey'
	size += len(d.Pubkey)

	return
}

// HashTreeRoot ssz hashes the DepositMessage object
func (d *DepositMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the DepositMessage object with a hasher
func (d *DepositMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Pubkey[:]))
		if byteLen > 48 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(d.Pubkey[:])
		hh.MerkleizeWithMixin(elemIndx, byteLen, (48+31)/32)
	}

	// Field (1) 'Credentials'
	hh.PutBytes(d.Credentials[:])

	// Field (2) 'Amount'
	hh.PutUint64(uint64(d.Amount))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the DepositMessage object
func (d *DepositMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}

// MarshalSSZ ssz marshals the Validator object
func (v *Validator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Validator object to a target array
func (v *Validator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkey'
	dst = append(dst, v.Pubkey[:]...)

	// Field (1) 'WithdrawalCredentials'
	dst = append(dst, v.WithdrawalCredentials[:]...)

	// Field (2) 'EffectiveBalance'
	dst = ssz.MarshalUint64(dst, uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	dst = ssz.MarshalBool(dst, v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.WithdrawableEpoch))

	return
}

// UnmarshalSSZ ssz unmarshals the Validator object
func (v *Validator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 121 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkey'
	copy(v.Pubkey[:], buf[0:48])

	// Field (1) 'WithdrawalCredentials'
	copy(v.WithdrawalCredentials[:], buf[48:80])

	// Field (2) 'EffectiveBalance'
	v.EffectiveBalance = primitives.Gwei(ssz.UnmarshallUint64(buf[80:88]))

	// Field (3) 'Slashed'
	v.Slashed = ssz.UnmarshalBool(buf[88:89])

	// Field (4) 'ActivationEligibilityEpoch'
	v.ActivationEligibilityEpoch = primitives.Epoch(ssz.UnmarshallUint64(buf[89:97]))

	// Field (5) 'ActivationEpoch'
	v.ActivationEpoch = primitives.Epoch(ssz.UnmarshallUint64(buf[97:105]))

	// Field (6) 'ExitEpoch'
	v.ExitEpoch = primitives.Epoch(ssz.UnmarshallUint64(buf[105:113]))

	// Field (7) 'WithdrawableEpoch'
	v.WithdrawableEpoch = primitives.Epoch(ssz.UnmarshallUint64(buf[113:121]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Validator object
func (v *Validator) SizeSSZ() (size int) {
	size = 121
	return
}

// HashTreeRoot ssz hashes the Validator object
func (v *Validator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Validator object with a hasher
func (v *Validator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	hh.PutBytes(v.Pubkey[:])

	// Field (1) 'WithdrawalCredentials'
	hh.PutBytes(v.WithdrawalCredentials[:])

	// Field (2) 'EffectiveBalance'
	hh.PutUint64(uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	hh.PutBool(v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	hh.PutUint64(uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	hh.PutUint64(uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	hh.PutUint64(uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	hh.PutUint64(uint64(v.WithdrawableEpoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Validator object
func (v *Validator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}